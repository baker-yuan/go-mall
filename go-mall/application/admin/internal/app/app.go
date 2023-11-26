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
	l := logger.New(cfg.Log.Level)

	// 初始化数据库
	conn, err := db.GetConn(cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Timeout, cfg.DB.DbName)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - db.GetConn: %w", err))
	}

	// gorm事务封装
	db.InitTransaction(conn)

	// gorm创建表
	if err := entity.Init(conn); err != nil {
		l.Fatal(fmt.Errorf("app - Run - entity.Init: %w", err))
	}
	// 全字段更新，初始化那些字段不更新，那些字段需要更新
	if err := repo.InitField(conn); err != nil {
		l.Fatal(fmt.Errorf("app - Run - repo.InitField: %w", err))
	}
	// oss url 前缀
	util.InitBaseUrl(cfg.Oss.BaseUrl)

	// 业务逻辑
	categoryUseCase := usecase.NewCategoryUseCase(
		repo.NewProductCategoryRepo(conn),
		repo.NewProductCategoryAttributeRelationRepo(conn),
		repo.NewProductRepo(conn),
	)

	brandUseCase := usecase.NewBrandUseCase(
		repo.NewBrandRepo(conn),
	)
	productAttributeCategoryUseCase := usecase.NewProductAttributeCategory(
		repo.NewProductAttributeCategoryRepo(conn),
	)
	productAttributeUseCase := usecase.NewProductAttribute(
		repo.NewProductAttributeRepo(conn),
	)
	skuStockUseCase := usecase.NewSkuStock(
		repo.NewSkuStockRepo(conn),
	)
	productUseCase := usecase.NewProduct(
		repo.NewProductRepo(conn),
		repo.NewBrandRepo(conn),
		repo.NewProductCategoryRepo(conn),
		repo.NewMemberPriceRepo(conn),
		repo.NewProductLadderRepo(conn),
		repo.NewProductFullReductionRepo(conn),
		repo.NewSkuStockRepo(conn),
		repo.NewProductAttributeValueRepo(conn),
		repo.NewSubjectProductRelationRepo(conn),
		repo.NewPrefrenceAreaProductRelationRepo(conn),
	)

	// grpc服务
	impl := grpcsrv.New(
		categoryUseCase,
		brandUseCase,
		productAttributeCategoryUseCase,
		productAttributeUseCase,
		productUseCase,
		skuStockUseCase,
	)
	if err := configGrpc(impl, cfg.HTTP.IP, cfg.HTTP.Port); err != nil {
		l.Fatal(fmt.Errorf("app - Run - configGrpc: %w", err))
	}

	// 监听关闭信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
		//case err = <-httpServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// 处理关闭
	//err = httpServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	//}

}

func configGrpc(impl pb.AdminApiServer, ip string, port uint32) error {
	var (
		addr = fmt.Sprintf("%s:%d", ip, port)
	)
	// 创建一个gRPC server对象
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ValidationInterceptor),
	)

	// 注册grpc服务
	pb.RegisterAdminApiServer(grpcServer, impl)

	// gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminApiHandlerFromEndpoint(context.Background(), gwmux, addr, dops); err != nil {
		return err
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
		return err
	}

	go func() {
		log.Fatalln(gwServer.Serve(lis)) // 启动HTTP服务
	}()

	return nil
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
