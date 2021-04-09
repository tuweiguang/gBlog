package db

import (
	"fmt"
	"gBlog/common/config"
	"gBlog/common/log"
	"github.com/jinzhu/gorm"
)

// 目前支持以下数据库
const (
	DB_Mysql = "mysql"
	DB_Redis = "redis"
	DB_IP    = "ip2region"
)

type SQL interface{}

type Base interface {
	init(dbI SQL) (db interface{})
}

func Init() {
	var base Base
	dbs := config.GetDBConfig()
	for _, db := range dbs {
		if db.DbType == DB_Mysql {
			base = &MySQL{}
		} else if db.DbType == DB_Redis {
			base = &Redis{}
		} else if db.DbType == DB_IP {
			base = &IP2Region{}
		}

		if base == nil {
			continue
		}
		instances[db.DbType] = base.init(db)
		log.GetLog().Info(fmt.Sprintf("%v init success", db.DbType))
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
		case DB_IP:
			ip, ok := v.(*GIP2Region)
			if !ok {
				continue
			}
			ip.Close()
		}
	}
}

var instances = make(map[string]interface{})

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

//获取ip
func GetIP() *GIP2Region {
	instance, ok := instances[DB_IP]
	if !ok {
		log.GetLog().Error("DB GetIP,no ip2region instance")
		panic("DB GetIP,no ip2region instance")
		return nil
	}

	ip, ok := instance.(*GIP2Region)
	if !ok {
		log.GetLog().Error("DB GetIP,instance is not *GIP2Region")
		panic("DB GetIP,instance is not *GIP2Region")
		return nil
	}
	return ip
}
