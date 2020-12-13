package admin

import (
	"gBlog/commom/sys"
	"gBlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const LIMIT = 10 //一页10条记录
const ONLINE = 1
const UNSALE = 2
const DELETE = 3

var Status = map[int]string{ONLINE: "在线", UNSALE: "下架", DELETE: "删除"}

type UserCtl struct{}

func (u *UserCtl) List(c *gin.Context) {
	page := c.GetInt("page")
	if page < 1 {
		page = 1
	}

	name := c.Query("name")
	status := c.GetInt("status")

	some := models.GetSomeUser((page-1)*LIMIT, LIMIT)
	c.HTML(http.StatusOK, "user-list.html", gin.H{
		"Status":     status,
		"Name":       name,
		"Data":       some,
		"Paginator":  models.GenPaginator(page, LIMIT, len(some)),
		"StatusText": Status,
	})
}

func (u *UserCtl) Welcome(c *gin.Context) {
	df, _ := sys.Df()
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"Df": df,
	})
}

func (u *UserCtl) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "user-add.html", nil)
}

func (u *UserCtl) Register(c *gin.Context) {
	if c.Request.Method == "POST" {
		var info RegisterInfo
		err := c.ShouldBind(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "填写参数错误",
			})
			return
		}

		// password == repassword
		if info.Password != info.RePassword {
			c.JSON(http.StatusOK, gin.H{
				"msg": "输入密码不正确",
			})
			return
		}

		//需要到数据库中找这个username是否存在
		if models.IsUserExists(info.Username) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "账户已存在",
			})
			return
		}

		//不存在则保存到数据库
		user := &models.User{
			Name:     info.Username,
			Password: info.Password,
			Email:    info.Email,
			Created:  time.Now(),
			Status:   0,
		}

		if models.SaveToDB(user) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "创建用户成功",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "创建用户失败",
		})
		return
	}
}
