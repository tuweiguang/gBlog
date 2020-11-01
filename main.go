package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("view/*")

	r.GET("/detail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "view/detail.html", gin.H{
			// template.HTML 让模板中的参数不要做HTML转义
			"data": template.HTML(""),
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})

	r.Run()
}
