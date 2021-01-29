package middleware

import (
	"gBlog/controllers/homepage"
	"github.com/gin-gonic/gin"
)

func Prepare() gin.HandlerFunc {
	comCtl := new(homepage.CommonCtl)
	return func(c *gin.Context) {
		comCtl.Archives(c)
		comCtl.Menu(c)
		comCtl.Keywords(c)

		c.Next()
	}
}

//人数
func StatisticsUV() gin.HandlerFunc {
	comCtl := new(homepage.CommonCtl)
	return func(c *gin.Context) {
		c.Next()

		comCtl.UV(c)
	}
}

//点击量
func StatisticsPV() gin.HandlerFunc {
	comCtl := new(homepage.CommonCtl)
	return func(c *gin.Context) {
		c.Next()

		comCtl.PV(c)
	}
}
