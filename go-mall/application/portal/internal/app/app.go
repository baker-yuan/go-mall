// Package app configures and runs application.
package app

import (
	"github.com/evrone/go-clean-template/config"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	//l := logger.New(cfg.Log.Level)
	//
	//// 初始化数据库
	//conn, err := db.GetConn(cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Timeout, cfg.DB.DbName)
	//if err != nil {
	//	customLog.Fatal(fmt.Errorf("app - Run - db.GetConn: %w", err))
	//}
	//
	//// gorm事务封装
	//db.InitTransaction(conn)
	//
	//// 业务逻辑
	//translationUseCase := usecase.New(
	//	repo.New(pg), // 依赖数据操作
	//	webapi.New(), // 依赖外部接口
	//)
	//
	//// http服务器
	//handler := gin.New()
	//v1.NewRouter(handler, l, translationUseCase)
	//httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	//
	//// 监听系统信号
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	//select {
	//case s := <-interrupt:
	//	l.Info("app - Run - signal: " + s.String())
	//case err = <-httpServer.Notify():
	//	l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	//}
	//
	//// 关闭http服务
	//err = httpServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	//}

}
