package admin

import (
	"gBlog/common"
	"gBlog/common/config"
	"gBlog/common/sys"
	"gBlog/common/util"
	"gBlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

const ONLINE = 1
const UNSALE = 2
const DELETE = 3

var Status = map[int]string{ONLINE: "在线", UNSALE: "下架", DELETE: "删除"}

type UserCtl struct{}

func (u *UserCtl) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	name := c.Query("name")
	status, _ := strconv.Atoi(c.Query("status"))

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
	var dates []string
	var pv []int
	var uv []int

	now := time.Now()
	for i := PV_UV_LIMIT; i > 0; i-- {
		dates = append(dates, now.AddDate(0, 0, -i).Format("20060102"))
	}

	for _, date := range dates {
		pv = append(pv, models.GetPVByDay(date))
		uv = append(uv, models.GetUVByDay(date))
	}

	dates = nil
	for i := PV_UV_LIMIT; i > 0; i-- {
		dates = append(dates, now.AddDate(0, 0, -i).Format("0102"))
	}

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"Df":   df,
		"PV":   pv,
		"UV":   uv,
		"Date": dates,
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

		pw := info.Password
		if config.GetAPPConfig().Env == common.EnvRelease {
			pw = util.PasswordMD5(info.Password, info.Username)
		}
		//不存在则保存到数据库
		user := &models.User{
			Name:     info.Username,
			Password: pw,
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
