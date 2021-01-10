package homepage

import (
	"gBlog/models"
	"github.com/gin-gonic/gin"
	"strings"
)

type CommonCtl struct{}

func (cc *CommonCtl) Archives(c *gin.Context) {

}

func (cc *CommonCtl) Menu(c *gin.Context) {

}

func (cc *CommonCtl) Keywords(c *gin.Context) {
	all := models.GetAllArticle()

	var tags []string
	for _, v := range all {
		tags = append(tags, strings.Split(strings.Replace(v.Tag, `，`, `,`, -1), `,`)...)
	}

	var tagsMap = make(map[string]int)

	for _, v := range tags {
		tagsMap[v] += 1
	}

	for k := range tagsMap {
		tagsMap[k] = tagsMap[k]/5 + 15
	}

	res := gin.H{}
	res["Tag"] = tagsMap
	//c.Set("Tag", tagsMap)

	count := len(all)
	var datetime = make(map[string]int64)
	var dateTimeKey []string
	for _, v := range all {
		k := v.CreatedAt.Format("2006-01")
		if datetime[k] == 0 {
			dateTimeKey = append(dateTimeKey, k)
		}
		datetime[k] = datetime[k] + 1
	}
	//c.Set("DateTime", datetime)
	//c.Set("DateTimeKey", dateTimeKey)
	//c.Set("DateCount", count)
	res["DateTime"] = datetime
	res["DateTimeKey"] = dateTimeKey
	res["DateCount"] = count
	c.Set("common", res)
}
