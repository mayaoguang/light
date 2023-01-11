package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	"google.golang.org/grpc"
	"light/internal/config"
	"light/pkg/logging"
	"light/pkg/mq"
	"light/pkg/mysql"
	"light/pkg/redis"
	"sync"
	"time"
)

var (
	err         error
	configPath  string
	ctx, cancel = context.WithCancel(context.Background())
	wg          = new(sync.WaitGroup)
	app         = iris.New()
	gServer     = grpc.NewServer()
)

func main() {
	defer func() {
		wg.Wait()
		logging.Sync()
	}()
	fmt.Println("light main")
}

func init() {
	flag.StringVar(&configPath, "config", "./configs/config.json", "配置文件路径以及文件名(必填)")
	flag.Parse()

	// 初始化配置
	config.Init(configPath)

	// 注册mysql
	if err = mysql.Init(config.Config.Mysql.Read, config.Config.Mysql.Write); err != nil {
		logging.Fatal("init mysql service err: " + err.Error())
	}

	// 注册redis
	if err = redis.Init(config.Config.Redis); err != nil {
		logging.Fatal("init redis service err: " + err.Error())
	}

	// 注册RabbitMQ
	if err = mq.Init(config.Config.Mq); err != nil {
		logging.Fatal("init rabbit mq err: " + err.Error())
	}

	// 优雅关闭程序
	iris.RegisterOnInterrupt(func() {
		wg.Add(1)
		defer wg.Done()
		cancel()
		time.Sleep(5 * time.Second)
		// 关闭所有主机
		gServer.Stop()
		_ = app.Shutdown(ctx)
	})
}
