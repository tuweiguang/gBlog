package middleware

import (
	"gBlog/pkg/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 非login接口
func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("sessionId")
		if err == nil {
			// 到本地或者redis里面去验证sessionId
			if status := session.NewMemoryMgr().CheckSession(cookie); status > session.SessionExist {
				if status == session.SessionExpire {
					// 删除session
					session.NewMemoryMgr().DelSession(cookie)
				}

				// sessionId不存在或者过期，需要重新登陆
				c.Abort()
				c.Redirect(http.StatusMovedPermanently, "/login")
				return
			}
			c.Next() //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		} else {
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/login")
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}

// https://www.cnblogs.com/wind-zhou/p/13114548.html
