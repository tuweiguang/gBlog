package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleCtl struct{}

func (a *ArticleCtl) List(c *gin.Context) {

	c.HTML(http.StatusOK, "article-list.html", nil)
}
