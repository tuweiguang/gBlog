package homepage

import (
	"fmt"
	"gBlog/controllers/admin"
	"gBlog/models"
	"gBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ContentCtl struct{}

func (con *ContentCtl) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1")) // 页数
	offset := (page - 1) * admin.LIMIT                   // 偏移量

	cate, _ := strconv.Atoi(c.Query("page"))
	//tag,_ := strconv.Atoi(c.Query("tag")) //根据tag获取
	//date,_ := strconv.Atoi(c.Query("date"))//根据日期获取

	var res gin.H
	res, ok := c.MustGet("common").(gin.H)
	if !ok {
		res = gin.H{}
	}
	fmt.Printf("====================> common:%v \n", res)
	res["CategoryID"] = cate
	articles := models.GetSomeArticle(offset, admin.LIMIT)

	var users []int64
	var categorys []int
	for _, v := range articles {
		if !utils.IsExistsElementInt64(users, v.UserId) {
			users = append(users, v.UserId)
		}

		if !utils.IsExistsElementInt(categorys, v.CategoryId) {
			categorys = append(categorys, v.CategoryId)
		}
	}
	someUser := models.GetSomeUserByIds(users)
	someCate := models.GetCategoryById(categorys)
	for i, v := range articles {
		for _, vv := range someUser {
			if v.UserId == vv.Id {
				articles[i].User = vv
				break
			}
		}

		for _, vv := range someCate {
			if v.CategoryId == vv.Id {
				articles[i].Category = vv
				break
			}
		}
	}

	res["Data"] = articles
	res["Paginator"] = models.GenPaginator(page, admin.LIMIT, len(articles))

	c.HTML(http.StatusOK, "front-list.html", res)
}
