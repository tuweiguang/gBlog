package db

import (
	"gBlog/commom/config"
	"gBlog/commom/log"
	"github.com/jinzhu/gorm"
)

// 目前支持以下数据库
const (
	DB_Mysql = "mysql"
	DB_Redis = "redis"
)

type SQL interface{}

type Base interface {
	init(dbI SQL) (db interface{})
}

func Init() {
	var base Base
	dbs := config.GetDBConfig()
	for _, db := range dbs {
		if db.DbType == "mysql" {
			base = &MySQL{}
		} else if db.DbType == "redis" {
			base = &Redis{}
		}

		if base == nil {
			continue
		}
		instances[db.DbType] = base.init(db)
	}
}

func Close() {
	for k, v := range instances {
		switch k {
		case DB_Mysql:
			mysql, ok := v.(*gorm.DB)
			if !ok {
				continue
			}
			mysql.Close()
		case DB_Redis:
			redis, ok := v.(*GRedis)
			if !ok {
				continue
			}
			redis.Close()
		}
	}
}

var instances map[string]interface{}

// 获取 mysql连接
func GetMySQL() *gorm.DB {
	instance, ok := instances[DB_Mysql]
	if !ok {
		log.GetLog().Error("DB GetDB,no mysql instance")
		panic("DB GetDB,no mysql instance")
		return nil
	}
	mysql, ok := instance.(*gorm.DB)
	if !ok {
		log.GetLog().Error("DB GetDB,instance is not *gorm.DB")
		panic("DB GetDB,instance is not *gorm.DB")
		return nil
	}
	return mysql
}

// 获取redis连接
func GetRedis() *GRedis {
	instance, ok := instances[DB_Redis]
	if !ok {
		log.GetLog().Error("DB GetRedis,no redis instance")
		panic("DB GetRedis,no redis instance")
		return nil
	}

	redis, ok := instance.(*GRedis)
	if !ok {
		log.GetLog().Error("DB GetRedis,instance is not *GRedis")
		panic("DB GetRedis,instance is not *GRedis")
		return nil
	}
	return redis
}
