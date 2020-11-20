package controllers

import (
	"gBlog/commom/config"
	"gBlog/commom/util"
	"gBlog/controllers/admin"
	"gBlog/middleware"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

var engine *gin.Engine

func init() {
	engine = gin.Default()

	// 设置模板函数
	engine.SetFuncMap(template.FuncMap{
		"IndexForOne":  util.IndexForOne,
		"IndexDecrOne": util.IndexDecrOne,
		"IndexAddOne":  util.IndexAddOne,
	})

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

	// 1.无需认证
	// /list.html
	// /detail/:id([0-9]+).html

	// 2.需要认证
	engine.Use(middleware.PrintSession())
	engine.Any("/login", admin.Login)

	adminCtl := engine.Group("/admin", middleware.Validate())
	{
		adminCtl.GET("/list", admin.List)

		adminCtl.GET("/welcome", admin.Welcome)

		adminCtl.GET("/user", func(context *gin.Context) {})
		adminCtl.GET("/user/add", func(context *gin.Context) {})

		adminCtl.GET("/article", func(context *gin.Context) {})
		adminCtl.GET("/article/edit", func(context *gin.Context) {})
		adminCtl.GET("/article/delete", func(context *gin.Context) {})
		adminCtl.GET("/article/update", func(context *gin.Context) {})
		adminCtl.GET("/article/add", func(context *gin.Context) {})
		adminCtl.GET("/article/top", func(context *gin.Context) {})
		adminCtl.GET("/article/get", func(context *gin.Context) {})

		adminCtl.GET("/cate", func(context *gin.Context) {})
		adminCtl.GET("/cate/add", func(context *gin.Context) {})
		adminCtl.GET("/cate/edit", func(context *gin.Context) {})
		adminCtl.GET("/cate/delete", func(context *gin.Context) {})
		adminCtl.GET("/cate/update", func(context *gin.Context) {})
	}

	engine.GET("/admin", middleware.Validate(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// 3.没有匹配到走下面
	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}

func DefaultServerRun() {
	go engine.Run(config.GetAPPConfig().HttpAddr)
}
