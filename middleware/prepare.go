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
