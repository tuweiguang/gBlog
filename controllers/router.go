package controllers

import (
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

	admin := engine.Group("/admin")
	{
		admin.GET("/list", func(c *gin.Context) {

		})
		admin.POST("/ssave", func(c *gin.Context) {

		})
	}
}

func DefaultServerRun() {
	go engine.Run(":8080")
}
