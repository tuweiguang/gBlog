package models

import (
	"gBlog/commom/db"
	"strconv"
)

const ArticleUriPV = "articleUriPV"
const AccessUV = "accessUV"

// 每个路由访问量
func AddArticlePV(uri string) int64 {
	return db.GetRedis().HIncrBy(ArticleUriPV, uri, 1).Val()
}

func GetArticlePV(uri string) int {
	pv, _ := db.GetRedis().HGet(ArticleUriPV, uri).Int()
	return pv
}

func GetAllPV() int {
	m := db.GetRedis().HGetAll(ArticleUriPV).Val()
	pv := 0
	for _, v := range m {
		count, _ := strconv.Atoi(v)
		pv += count
	}
	return pv
}

func IncrUV() {
	db.GetRedis().Incr(AccessUV)
}

func GetUV() int {
	uv, _ := db.GetRedis().Get(AccessUV).Int()
	return uv
}
