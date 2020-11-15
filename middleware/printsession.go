package middleware

import (
	"gBlog/pkg/session"
	"github.com/gin-gonic/gin"
)

func PrintSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		session.NewMemoryMgr().PrintSession()
	}
}
