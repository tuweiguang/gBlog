package main

import (
	"flag"
	"fmt"
	"gBlog/common/config"
	"gBlog/common/db"
	"gBlog/common/log"
	"gBlog/common/monitor"
	"gBlog/controllers"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// go run main.go -conf ./conf
func main() {
	flag.Parse()

	config.Init() // flag: -conf ./conf 指定配置文件目录
	log.Init()
	db.Init()
	monitor.Init(config.GetMonitorConfig().DumpPath)

	// 开启pprof，监听请求
	go func() {
		if err := http.ListenAndServe(config.GetAPPConfig().PProfAddr, nil); err != nil {
			log.GetLog().Error(fmt.Sprintf("start pprof failed on %s\n", config.GetAPPConfig().PProfAddr))
		}
	}()

	controllers.DefaultServerRun()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			controllers.Shutdown()
			db.Close()
			log.GetLog().Sync()
			time.Sleep(time.Second)
			panic("service is closing!")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
