package admin

import (
	"fmt"
	"gBlog/commom/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCtl struct{}

func Login(c *gin.Context) {
	log.GetLog().Info("request /login")
	username := c.Query("username")
	password := c.Query("password")
	fmt.Printf("=============> username:%v,password:%v", username, password)
	if len(username) > 0 && len(password) > 0 {
		// 设置session
		// value不能为1??
		c.SetCookie("gBlog_cookie", "01", 3600, "/", "localhost", false, true)
		// 重定向
		c.Redirect(http.StatusMovedPermanently, "/admin")
	}

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (u *UserCtl) List(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (u *UserCtl) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{})
}
