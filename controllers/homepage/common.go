package homepage

import (
	"fmt"
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

//点击量
func (cc *CommonCtl) PV(c *gin.Context) {
	// 获取uri
	uri := c.FullPath()

	//需要在nginx配置如下:
	//location /go/ {
	//        proxy_set_header X-Forward-For $remote_addr;
	//        proxy_set_header X-real-ip $remote_addr;
	//        proxy_pass http://127.0.0.1:8080/go/;
	//}
	ip := c.ClientIP()

	models.CreateAccessLog(ip, "", uri)
	models.AddPV(uri)
	fmt.Printf("uri:%v ip:%v\n", uri, ip)
}

//人数(根据cookie来判断)
func (cc *CommonCtl) UV(c *gin.Context) {

}
