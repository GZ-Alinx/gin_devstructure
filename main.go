package main

import (
	"context"
	"flag"
	"fmt"
	"gin_web/controller"
	"gin_web/dao/mysqls"
	"gin_web/dao/redis"
	"gin_web/logger"
	"gin_web/pkg/snowflake"
	"gin_web/routes"
	"gin_web/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// Go web开发通用的脚手架模板

func main() {

	// 加载配置,使用flag进行参数指定
	var fconfig string
	flag.StringVar(&fconfig, "f", "./conf/config.yaml", "指定配置文件") // 设置参数
	// fconfig := flag.String("f", "./conf/config.yaml", "指定配置文件")  引用：  *fconfig
	flag.Parse() // 解析参数
	config := make([]string, 10)
	if err := settings.Init(fconfig); err != nil {
		fmt.Printf("init settings failed, err : %v \n", err)
		return
	}

	// 初始化日志
	defer zap.L().Sync() // 将缓冲区的日志追加到日志文件中
	if err := logger.Init(settings.Conf.LogsConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err : %v \n", err)
		return
	}
	zap.L().Debug("Logger load success ...")
	config = append(config, "Logger loads success ...")

	// 初始化MySQL连接
	if err := mysqls.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err : %v \n", err)
		return
	}
	defer mysqls.Close()
	config = append(config, "MySQL loads success ...")

	// 初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err : %v \n", err)
		return
	}
	defer mysqls.Close()
	config = append(config, "Redis loads success ...")

	// 通过雪花算法获取用户ID
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := routes.Setup(settings.Conf.Mode)
	config = append(config, "Router loads success ...")
	for _, v := range config {
		fmt.Println(v)
	}

	// 初始化gin框架内置校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Println("init validator trans failed , err ", err)
		return
	}

	// 启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.AppPort),
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("start Faild", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)                      // 创建一个信号通道，对通道中的信号进行捕获
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 捕获信号
	<-quit                                               // 阻塞， 当接收到上述两种信号时才会往下执行
	zap.L().Info("shutdown Server ...")

	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务， 超过5秒就超时退出）
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server Exiting")

}
