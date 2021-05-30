package models

import (
	"fmt"
	"gBlog/common/db"
	"strconv"
)

const ArticleUriPV = "articleUriPV"
const AccessUV = "accessUV"

const DailyPV = "dailyPV:%v"
const DailyUV = "dailyUV:%v"

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

// 记录每日PV量(异步)
func AddDailyPV(date string) {
	db.GetRedis().Incr(fmt.Sprintf(DailyPV, date))
}

func GetPVByDay(date string) int {
	pv, _ := db.GetRedis().Get(fmt.Sprintf(DailyPV, date)).Int()
	return pv
}

// 记录每日UV量(异步)
func AddDaliyUV(date string) {
	db.GetRedis().Incr(fmt.Sprintf(DailyUV, date))
}

func GetUVByDay(date string) int {
	uv, _ := db.GetRedis().Get(fmt.Sprintf(DailyUV, date)).Int()
	return uv
}
