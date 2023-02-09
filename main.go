package main

import (
	"context"
	"fmt"
	"log"
	"master-gin/controllers"
	"master-gin/dao/mysql"

	// "master-gin/dao/redis"
	"master-gin/logger"
	"master-gin/pkg/snowflake"
	"master-gin/router"
	"master-gin/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err %v\n", err)
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err %v\n", err)
	}
	// NOTE: force to  write memory in cache to disk
	defer zap.L().Sync()
	zap.L().Debug("logger init success!")

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err %v\n", err)
	}
	defer mysql.Close()

	// if err := redis.Init(settings.Conf.RedisConfig); err != nil {
	// 	fmt.Printf("init redis failed, err %v\n", err)
	// }
	// defer redis.Close()

	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	r := router.Setup()

	if err := controllers.InitTrans("en"); err != nil {
		fmt.Printf("init validator trans failed %v\n", err)
		return
	}

	/**
	*	graceful shutdown
	 */

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
