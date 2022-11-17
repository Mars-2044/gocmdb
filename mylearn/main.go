package main 

import (
	"context"
	"fmt"
	"log"
	"mylearn/dao/mysql"
	"mylearn/dao/redis"
	"mylearn/router"
	"mylearn/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// go web 开发较通用的脚手架模板

func main() {
	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err :%v\n", err)
		return
	}

	// 2. 初始化日志
	//if err := logger.Init(); err != nil {
	//	fmt.Printf("init logger failed, err :%v\n", err)
	//	return
	//}

	// 3. 初始化Mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err :%v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化Redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("init mysql failed, err :%v\n", err)
		return
	}
	defer redis.Close()

	// 5. 注册路由
	r := router.Setup()
	// 6. 启动服务(优雅关机)

	srv := &http.Server {
		//Addr:    fmt.Sprintf(":%d", viper.GetString("app.port")),
    Addr: ":8081",
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit
	log.Println("shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("server shutdown", err)
	}

	log.Println("Server exiting")
}
