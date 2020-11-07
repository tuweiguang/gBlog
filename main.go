package main

import (
	"flag"
	"fmt"
	"gBlog/commom/config"
	"gBlog/commom/db"
	"gBlog/commom/log"
	"gBlog/controllers"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()

	config.Init() // flag: -conf ./conf 指定配置文件目录
	log.Init()
	db.Init()

	controllers.DefaultServerRun()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("gBlog exit")
			db.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
