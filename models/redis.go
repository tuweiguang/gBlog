package models

import "gBlog/commom/db"

const ArticleUriPV = "articleUriPV"

// 每个路由访问量
func AddPV(uri string) {
	db.GetRedis().HIncrBy(ArticleUriPV, uri, 1)
}
