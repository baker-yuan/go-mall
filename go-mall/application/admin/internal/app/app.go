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

	config "github.com/baker-yuan/go-mall/application/admin/config"
	"github.com/baker-yuan/go-mall/application/admin/internal/controller/grpcsrv"
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/internal/usecase"
	"github.com/baker-yuan/go-mall/application/admin/internal/usecase/repo"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	"github.com/baker-yuan/go-mall/application/admin/pkg/interceptor"
	"github.com/baker-yuan/go-mall/application/admin/pkg/logger"
	"github.com/baker-yuan/go-mall/application/admin/pkg/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Run 启动服务
func Run(cfg *config.Config) {
	// 日志
	customLog := logger.New(cfg.Log.Level)

	// 初始化数据库
	conn, err := db.GetConn(cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Timeout, cfg.DB.DbName)
	if err != nil {
		customLog.Fatal(fmt.Errorf("app - Run - db.GetConn: %w", err))
	}

	// gorm事务封装
	db.InitTransaction(conn)

	// gorm创建表
	if err := entity.Init(conn); err != nil {
		customLog.Fatal(fmt.Errorf("app - Run - entity.Init: %w", err))
	}
	// 全字段更新，初始化那些字段不更新，那些字段需要更新
	if err := repo.InitField(conn); err != nil {
		customLog.Fatal(fmt.Errorf("app - Run - repo.InitField: %w", err))
	}
	// oss url 前缀
	util.InitBaseUrl(cfg.Oss.BaseUrl)

	var (
		productCategoryRepo                  = repo.NewProductCategoryRepo(conn)
		productCategoryAttributeRelationRepo = repo.NewProductCategoryAttributeRelationRepo(conn)
		productRepo                          = repo.NewProductRepo(conn)
		brandRepo                            = repo.NewBrandRepo(conn)
		productAttributeCategoryRepo         = repo.NewProductAttributeCategoryRepo(conn)
		productAttributeRepo                 = repo.NewProductAttributeRepo(conn)
		skuStockRepo                         = repo.NewSkuStockRepo(conn)
		memberPriceRepo                      = repo.NewMemberPriceRepo(conn)
		productLadderRepo                    = repo.NewProductLadderRepo(conn)
		productFullReductionRepo             = repo.NewProductFullReductionRepo(conn)
		productAttributeValueRepo            = repo.NewProductAttributeValueRepo(conn)
		subjectProductRelationRepo           = repo.NewSubjectProductRelationRepo(conn)
		prefrenceAreaProductRelationRepo     = repo.NewPrefrenceAreaProductRelationRepo(conn)
		subjectRepo                          = repo.NewSubjectRepo(conn)
		prefrenceAreaRepo                    = repo.NewPrefrenceAreaRepo(conn)
		//
		orderReturnReasonRepo   = repo.NewOrderReturnReasonRepo(conn)
		orderRepo               = repo.NewOrderRepo(conn)
		orderItemRepo           = repo.NewOrderItemRepo(conn)
		orderOperateHistoryRepo = repo.NewOrderOperateHistoryRepo(conn)
		orderReturnApplyRepo    = repo.NewOrderReturnApplyRepo(conn)
		companyAddressRepo      = repo.NewCompanyAddressRepo(conn)
	)

	// 业务逻辑
	categoryUseCase := usecase.NewCategoryUseCase(
		productCategoryRepo,
		productCategoryAttributeRelationRepo,
		productRepo,
	)
	brandUseCase := usecase.NewBrandUseCase(
		brandRepo,
	)
	productAttributeCategoryUseCase := usecase.NewProductAttributeCategory(
		productAttributeCategoryRepo,
	)
	productAttributeUseCase := usecase.NewProductAttribute(
		productAttributeRepo,
	)
	skuStockUseCase := usecase.NewSkuStock(
		skuStockRepo,
	)
	productUseCase := usecase.NewProduct(
		productRepo,
		brandRepo,
		productCategoryRepo,
		memberPriceRepo,
		productLadderRepo,
		productFullReductionRepo,
		skuStockRepo,
		productAttributeValueRepo,
		subjectProductRelationRepo,
		prefrenceAreaProductRelationRepo,
	)
	subjectUseCase := usecase.NewSubject(subjectRepo)
	prefrenceAreaUseCase := usecase.NewPrefrenceArea(prefrenceAreaRepo)

	orderReturnReasonUseCase := usecase.NewOrderReturnReason(orderReturnReasonRepo)
	orderUseCase := usecase.NewOrder(orderRepo, orderItemRepo, orderOperateHistoryRepo)
	orderReturnApplyUseCase := usecase.NewOrderReturnApply(orderReturnApplyRepo, companyAddressRepo)
	companyAddressUseCase := usecase.NewCompanyAddress(companyAddressRepo)

	// grpc服务
	grpcSrvImpl := grpcsrv.New(
		categoryUseCase,
		brandUseCase,
		productAttributeCategoryUseCase,
		productAttributeUseCase,
		productUseCase,
		skuStockUseCase,
		subjectUseCase,
		prefrenceAreaUseCase,
		//
		orderReturnReasonUseCase,
		orderUseCase,
		orderReturnApplyUseCase,
		companyAddressUseCase,
	)
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

func configGrpc(customLog *logger.Logger, grpcSrvImpl grpcsrv.AdminApi, ip string, port uint32) (*grpc.Server, error) {
	var (
		addr = fmt.Sprintf("%s:%d", ip, port)
	)
	// 创建一个gRPC server对象
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ValidationInterceptor),
	)

	// 注册grpc服务
	pb.RegisterCmsAdminApiServer(grpcServer, grpcSrvImpl)
	pb.RegisterOmsAdminApiServer(grpcServer, grpcSrvImpl)

	// gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterCmsAdminApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}
	if err := pb.RegisterOmsAdminApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// 定义HTTP server配置
	gwServer := &http.Server{
		Addr:    addr,
		Handler: grpcHandlerFunc(grpcServer, mux), // 请求的统一入口
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
