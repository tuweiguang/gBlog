package admin

import (
	"fmt"
	"gBlog/commom/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCtl struct{}

func (u *UserCtl) Login(c *gin.Context) {
	log.GetLog().Info("request /admin/login")
	username := c.Query("username")
	password := c.Query("password")
	fmt.Printf("=============> username:%v,password:%v", username, password)
	if len(username) > 0 && len(password) > 0 {
		// 重定向
		c.Redirect(http.StatusMovedPermanently, "/admin")
	}

	c.HTML(http.StatusOK, "login.html", gin.H{})
}
