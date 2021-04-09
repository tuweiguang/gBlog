package middleware

import (
	"fmt"
	"gBlog/common/config"
	"gBlog/common/log"
	"gBlog/controllers/homepage"
	"gBlog/pkg/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 非login接口
func Validate(e *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("sessionId")
		if err == nil {
			// 到本地或者redis里面去验证sessionId
			if status := session.NewMemoryMgr().CheckSession(cookie); status > session.SessionExist {
				url := c.Request.URL.String()
				log.GetLog().Info(fmt.Sprintf("url = %v, invalid cookie!", url))
				if status == session.SessionExpire {
					// 删除session
					session.NewMemoryMgr().DelSession(cookie)
				}

				// sessionId不存在或者过期，需要重新登陆
				c.Abort()
				c.Redirect(http.StatusMovedPermanently, "/login")
				return
			}
			url := c.Request.URL.String()
			log.GetLog().Info(fmt.Sprintf("url = %v, Valid cookie!", url))
			c.Next() //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		} else {
			url := c.Request.URL.String()
			log.GetLog().Info(fmt.Sprintf("url = %v, No cookie!", url))
			c.Abort()

			c.Redirect(http.StatusMovedPermanently, "/login")
			//c.Request.URL.Path = "/login"
			//e.HandleContext(c)
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}

// https://www.cnblogs.com/wind-zhou/p/13114548.html

func StatisticsPVAndUV(e *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		comCtl := new(homepage.CommonCtl)
		cookie, err := c.Cookie("sessionId")
		if err == nil {
			// 到本地或者redis里面去验证sessionId
			if status := session.NewMemoryMgr().CheckSession(cookie); status > session.SessionExist {
				url := c.Request.URL.String()
				log.GetLog().Info(fmt.Sprintf("url = %v, invalid cookie!", url))
				if status == session.SessionExpire {
					// 删除session
					session.NewMemoryMgr().DelSession(cookie)
				}

				//设置cookie
				sessionId := session.NewMemoryMgr().CreateSessoin()
				c.SetCookie(config.GetSessionConfig().Name, sessionId, config.GetSessionConfig().Expire,
					config.GetSessionConfig().Path, config.GetSessionConfig().Domain, config.GetSessionConfig().Secure,
					config.GetSessionConfig().HttpOnly)

				//增加UV
				comCtl.UV(c)
			}
		} else {
			//设置cookie
			sessionId := session.NewMemoryMgr().CreateSessoin()
			c.SetCookie(config.GetSessionConfig().Name, sessionId, config.GetSessionConfig().Expire,
				config.GetSessionConfig().Path, config.GetSessionConfig().Domain, config.GetSessionConfig().Secure,
				config.GetSessionConfig().HttpOnly)
			comCtl.UV(c)
		}
		comCtl.PV(c)

		c.Next()
	}
}
