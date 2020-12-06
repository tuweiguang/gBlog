package admin

import (
	"fmt"
	"gBlog/commom/config"
	"gBlog/models"
	"gBlog/pkg/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginCrl struct{}

type LoginInfo struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (l *LoginCrl) Login(c *gin.Context) {
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
		u := models.GetUser(info.Username)
		if u.Password == info.Password {
			sessionId, err := c.Cookie("sessionId")
			if err != nil {
				// 第一次来，没有sessionid，-->给用户建一个sessiondata，分配一个sessionid
				sessionId = session.NewMemoryMgr().CreateSessoin()

				// 设置session
				// value不能为1??
				// maxAge最好和session保存时间一样
				// httpOnly:true  js 脚本不能获取 cookie，可以防止跨站攻击，增加爬虫程序的难度
				// domain: 要注意这个参数，设置什么就要在浏览器写什么
				// 大坑：在浏览器必须输入http://localhost:8080/xxx 不能是http://127.0.0.1:8080/xxx,不然登陆返回cookie将在下次请求的时候不会携带，导致登陆不上
				c.SetCookie("sessionId", sessionId, config.GetAPPConfig().SessionExpire, "/", "localhost", false, true)
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
