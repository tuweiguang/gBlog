package controllers

import (
	"gBlog/commom/config"
	"gBlog/controllers/admin"
	"gBlog/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var engine *gin.Engine

func init() {
	engine = gin.Default()

	//渲染模板
	engine.LoadHTMLGlob("view/*")

	//配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	engine.StaticFS("/static", http.Dir("./static"))

	//engine.GET("/detail", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "view/detail.html", gin.H{
	//		// template.HTML 让模板中的参数不要做HTML转义
	//		"data": template.HTML(""),
	//	})
	//})

	engine.GET("/login", admin.Login)

	userCtl := new(admin.UserCtl)
	admin := engine.Group("/admin", middleware.Validate())
	{
		admin.GET("/list", userCtl.List)

		admin.GET("welcome", userCtl.Welcome)
	}

	engine.GET("/admin", middleware.Validate(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	//没有匹配到走下面
	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}

func DefaultServerRun() {
	go engine.Run(config.GetAPPConfig().HttpAddr)
}
