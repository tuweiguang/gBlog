package admin

import (
	"fmt"
	"gBlog/pkg/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		var info LoginInfo
		err := c.ShouldBind(&info)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "账号或密码不能为空!",
			})
			return
		}
		fmt.Printf("=============> username:%v,password:%v", info.Username, info.Password)
		// 应该去数据库去验证
		if len(info.Username) > 0 && len(info.Password) > 0 {
			sessionId, err := c.Cookie("sessionId")
			if err != nil {
				// 第一次来，没有sessionid，-->给用户建一个sessiondata，分配一个sessionid
				sessionId = session.NewMemoryMgr().CreateSessoin()

				// 设置session
				// value不能为1??
				// maxAge最好和session保存时间一样
				c.SetCookie("sessionId", sessionId, 60, "/", "localhost", false, true)
			} else {
				// 到本地或者redis里面去验证sessionId
				if status := session.NewMemoryMgr().CheckSession(sessionId); status > session.SessionExist {
					if status == session.SessionExpire {
						// 删除session
						session.NewMemoryMgr().DelSession(sessionId)
					}

					// sessionId不存在或者过期，需要重新登陆
					c.HTML(http.StatusOK, "login.html", gin.H{
						"err": "账号登陆过期，请重新登陆!",
					})
				}
			}

			// 重定向
			c.Redirect(http.StatusMovedPermanently, "/admin")
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err": "账号或密码不正确，请重新输入!",
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
}

func List(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{})
}
