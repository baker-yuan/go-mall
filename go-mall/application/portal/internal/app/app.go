// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/baker-yuan/go-mall/application/portal/config"
	"github.com/baker-yuan/go-mall/application/portal/internal/controller/grpcsrv"
	"github.com/baker-yuan/go-mall/application/portal/internal/usecase"
	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/repo"
	"github.com/baker-yuan/go-mall/common/db"
	"github.com/baker-yuan/go-mall/common/interceptor"
	"github.com/baker-yuan/go-mall/common/logger"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	customLog := logger.New(cfg.Log.Level)

	// 初始化数据库
	conn, err := db.GetConn(cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Timeout, cfg.DB.DbName)
	if err != nil {
		customLog.Fatal(fmt.Errorf("app - Run - db.GetConn: %w", err))
	}

	// gorm事务封装
	db.InitTransaction(conn)

	// oss url 前缀
	util.InitBaseUrl(cfg.Oss.BaseUrl)

	var (
		productCategoryRepo       = repo.NewProductCategoryRepo(conn)
		productRepo               = repo.NewProductRepo(conn)
		brandRepo                 = repo.NewBrandRepo(conn)
		productAttributeRepo      = repo.NewProductAttributeRepo(conn)
		productAttributeValueRepo = repo.NewProductAttributeValueRepo(conn)
		skuStockRepo              = repo.NewSkuStockRepo(conn)
	)
	homeUseCase := usecase.NewHome(productCategoryRepo)
	productUseCase := usecase.NewProduct(
		productRepo,
		brandRepo,
		productAttributeRepo,
		productAttributeValueRepo,
		skuStockRepo,
	)

	// grpc服务
	grpcSrvImpl := grpcsrv.New(homeUseCase, productUseCase)
	grpcServer, err := configGrpc(customLog, grpcSrvImpl, cfg.HTTP.IP, cfg.HTTP.Port)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - configGrpc: %w", err))
	}

	// 打印当前进程的 ID
	customLog.Info("project started with pid %d", os.Getpid())

	// 监听关闭信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		customLog.Info("app - Run - signal: " + s.String())
	}

	// grpc优雅关闭，5s内必须完成关闭
	gracefulStopWithTimeout(customLog, grpcServer, 5*time.Second)
}

// gracefulStopWithTimeout 尝试在给定的超时时间内优雅地关闭 gRPC 服务器。
// 如果服务器在超时时间内成功关闭，那么函数会打印一条信息并返回。
// 如果超时时间到了但服务器还没有关闭，那么函数会强制关闭服务器并打印一条错误信息。
func gracefulStopWithTimeout(customLog *logger.Logger, grpcServer *grpc.Server, timeout time.Duration) {
	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 创建一个通道，用于接收服务器关闭的信号
	ch := make(chan struct{})
	go func() {
		// 在一个新的 goroutine 中优雅地关闭服务器
		grpcServer.GracefulStop()
		// 当服务器关闭时，关闭通道
		close(ch)
	}()

	select {
	case <-ch:
		// 如果从通道接收到了信号，说明服务器已经成功关闭
		customLog.Info("gRPC server stopped gracefully")
	case <-ctx.Done():
		// 如果上下文超时，说明服务器没有在给定的时间内关闭，此时强制关闭服务器
		customLog.Error("gRPC server stop timeout, force stopping")
		grpcServer.Stop()
	}
}

// customHTTPError 自定义错误处理函数
func customHTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	// err := status.New(codes.InvalidArgument, "无效的参数").Err()

	var (
		grpcCode   = codes.Internal
		message    = ""
		defaultErr = CustomError{
			HTTPStatus: http.StatusInternalServerError,
			ErrorCode:  int32(grpcCode),
			Message:    message,
		}
	)
	// 检查错误类型
	st, ok := status.FromError(err)
	if ok {
		grpcCode = st.Code()
		message = st.Message()
	} else {
		grpcCode = codes.Internal
		message = err.Error()
	}

	// 查找自定义错误映射，如果没有找到，则使用默认的映射
	customErr, ok := customErrorMap[grpcCode]
	if !ok {
		customErr = defaultErr
	}

	// 设置响应的Content-Type和状态码
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(customErr.HTTPStatus)

	// 创建包含code和message的JSON响应体
	responseBody := struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	}{
		Code:    customErr.ErrorCode,
		Message: customErr.Message,
	}

	// 发送JSON响应
	if err := marshaler.NewEncoder(w).Encode(responseBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"code": 500, "message": "failed to marshal error message"}`))
	}
}

func configGrpc(customLog *logger.Logger, grpcSrvImpl grpcsrv.PortalApi, ip string, port uint32) (*grpc.Server, error) {
	var (
		addr = fmt.Sprintf("%s:%d", ip, port)
	)
	// 创建一个gRPC server对象
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ChainUnaryServerInterceptors(
			interceptor.ValidationInterceptor,
			interceptor.PanicRecoveryInterceptor,
			//interceptor.CorsInterceptor()
		),
		),
	)

	// 注册grpc服务
	pb.RegisterPortalHomeApiServer(grpcServer, grpcSrvImpl)
	pb.RegisterPortalProductApiServer(grpcServer, grpcSrvImpl)

	// gRPC-Gateway mux
	gwmux := runtime.NewServeMux(runtime.WithErrorHandler(customHTTPError))
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterPortalHomeApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}
	if err := pb.RegisterPortalProductApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// cors处理器
	corsWrapper := cors.AllowAll().Handler(mux)
	// 统一返回值处理
	responseWrapper := interceptor.WrapResponseMiddleware(corsWrapper)

	// 定义HTTP server配置
	gwServer := &http.Server{
		Addr:    addr,
		Handler: grpcHandlerFunc(grpcServer, responseWrapper), // 请求的统一入口
	}

	// tpc监听
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	go func() {
		customLog.Fatal(gwServer.Serve(lis)) // 启动HTTP服务
	}()

	return grpcServer, nil
}

// grpcHandlerFunc 将gRPC请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
