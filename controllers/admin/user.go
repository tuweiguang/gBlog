package admin

import (
	"gBlog/commom/sys"
	"gBlog/models"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

const LIMIT = 10 //一页10条记录
const ONLINE = 1
const UNSALE = 2
const DELETE = 3

var Status = map[int]string{ONLINE: "在线", UNSALE: "下架", DELETE: "删除"}

func List(c *gin.Context) {
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
		"Paginator":  GenPaginator(page, LIMIT, len(some)),
		"StatusText": Status,
	})
}

func Welcome(c *gin.Context) {
	df, _ := sys.Df()
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"Df": df,
	})
}

type Paginator struct {
	CurrentPage int `json:"currentPage"` //当前页
	PageSize    int `json:"pageSize"`    //每页数量
	TotalPage   int `json:"totalPage"`   //总页数
	TotalCount  int `json:"totalCount"`  //总数量
}

func GenPaginator(page, limit, count int) Paginator {
	var paginator Paginator
	paginator.TotalCount = count
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(limit)))
	paginator.PageSize = limit
	paginator.CurrentPage = page
	return paginator
}
