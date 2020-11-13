package middleware

import (
	"fmt"
	"gBlog/commom/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("gBlog_cookie")
		log.GetLog().Info(fmt.Sprintf("=====================> 失败, cookie:%v", cookie))
		if err == nil && cookie == "01" {
			log.GetLog().Info("=====================> 成功")
			c.Next() //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		} else {
			log.GetLog().Info("=====================> 失败")
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/login")
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}

// https://www.cnblogs.com/wind-zhou/p/13114548.html
