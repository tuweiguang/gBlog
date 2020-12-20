package controllers

import (
	"context"
	"fmt"
	"gBlog/commom/config"
	"gBlog/commom/log"
	"gBlog/commom/util"
	"gBlog/controllers/admin"
	"gBlog/middleware"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

var e *gin.Engine
var srv *http.Server

func init() {
	e = gin.Default()

	// 设置模板函数
	e.SetFuncMap(template.FuncMap{
		"IndexForOne":  util.IndexForOne,
		"IndexDecrOne": util.IndexDecrOne,
		"IndexAddOne":  util.IndexAddOne,
	})

	//渲染模板
	e.LoadHTMLGlob("view/*")

	//配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	e.StaticFS("/static", http.Dir("./static"))

	// 1.无需认证
	noAuthenticationRouter()
	// 2.需要认证
	authenticationRouter()
}

func DefaultServerRun() {
	srv = &http.Server{
		Addr:    config.GetAPPConfig().HttpAddr,
		Handler: e,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.GetLog().Fatal(fmt.Sprintf("DefaultServerRun listen error:%v", err))
		}
		log.GetLog().Info("DefaultServerRun init success")
	}()
}

// grace
func Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.GetLog().Fatal(fmt.Sprintf("Shutdown error:%v", err))
	}
	log.GetLog().Info("Server exiting!!!")
}

func authenticationRouter() {
	user := new(admin.UserCtl)
	article := new(admin.ArticleCtl)
	adminCtl := e.Group("/admin", middleware.Validate(e), middleware.PrintSession())
	{
		adminCtl.GET("/welcome", user.Welcome)

		adminCtl.GET("/user", func(context *gin.Context) {})
		adminCtl.GET("/user/add", user.Add)
		adminCtl.GET("/user/list", user.List)
		adminCtl.POST("/user/register", user.Register)

		adminCtl.GET("/article", func(context *gin.Context) {})
		adminCtl.GET("/article/list", article.List)
		adminCtl.GET("/article/edit", func(context *gin.Context) {})
		adminCtl.GET("/article/delete", func(context *gin.Context) {})
		adminCtl.POST("/article/upload", article.Upload)
		adminCtl.GET("/article/add", article.Add)
		adminCtl.GET("/article/top", func(context *gin.Context) {})
		adminCtl.GET("/article/get", func(context *gin.Context) {})

		adminCtl.GET("/cate", func(context *gin.Context) {})
		adminCtl.GET("/cate/add", func(context *gin.Context) {})
		adminCtl.GET("/cate/edit", func(context *gin.Context) {})
		adminCtl.GET("/cate/delete", func(context *gin.Context) {})
		adminCtl.GET("/cate/update", func(context *gin.Context) {})
	}
}

func noAuthenticationRouter() {
	login := new(admin.LoginCrl)
	e.Any("/login", middleware.PrintSession(), login.Login)
	e.GET("/admin", middleware.Validate(e), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	e.GET("/list.html", func(context *gin.Context) {
		//	c.HTML(http.StatusOK, "view/detail.html", gin.H{
		//		// template.HTML 让模板中的参数不要做HTML转义
		//		"data": template.HTML(""),
		//	})
	})
	e.GET("/detail/:id([0-9]+).html", func(context *gin.Context) {
		id := context.Param("id([0-9]+).html")
		fmt.Printf("==========> %v,%v", context.Request.URL, id)
		context.JSON(http.StatusOK, nil)
	})

	// 没有匹配到走下面
	e.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}
