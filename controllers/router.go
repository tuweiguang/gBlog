package controllers

import (
	"gBlog/commom/config"
	"gBlog/controllers/admin"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

var engine *gin.Engine

func init() {
	engine = gin.Default()
	engine.LoadHTMLGlob("view/*")
	engine.GET("/detail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "view/detail.html", gin.H{
			// template.HTML 让模板中的参数不要做HTML转义
			"data": template.HTML(""),
		})
	})

	userCtl := new(admin.UserCtl)
	admin := engine.Group("/admin")
	{
		admin.GET("/login", userCtl.Login)
	}

	engine.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

func DefaultServerRun() {
	go engine.Run(config.GetAPPConfig().HttpAddr)
}
