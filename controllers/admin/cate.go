package admin

import (
	"gBlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CateCtl struct{}

func (cate *CateCtl) List(c *gin.Context) {
	all := models.GetAllCategory()

	res := gin.H{}
	res["Category"] = all
	c.HTML(http.StatusOK, "cate.html", res)
}

func (cate *CateCtl) Add(c *gin.Context) {
	if c.Request.Method == "GET" {
		all := models.GetAllCategory()
		res := gin.H{}
		res["Category"] = all
		c.HTML(http.StatusOK, "cate-add.html", res)
	} else if c.Request.Method == "POST" {
		var info AddCate
		err := c.ShouldBind(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "参数错误!",
			})
			return
		}
		cate := &models.Category{
			Pid:  info.CateId,
			Name: info.CateName,
			Sort: info.Sort,
		}
		err = models.AddCategory(cate)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "添加失败！",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "添加成功！",
		})
	}
}
