package db

import (
	"fmt"
	"gBlog/commom/config"
	"gBlog/commom/log"
	"github.com/jinzhu/gorm"
)

// 目前只支持以下数据库
const (
	Mysql = "mysql"
)

type SQL interface{}

type Base interface {
	init(dbI SQL) (db *gorm.DB)
}

func Init() {
	var sql Base
	dbs := config.GetDBConfig()
	for _, db := range dbs {
		if db.DbType == "mysql" {
			sql = &MySQL{}
		}

		instances[db.DbType] = sql.init(db)
	}
}

func Close() {
	for _, v := range instances {
		_ = v.Close()
	}
}

var instances map[string]*gorm.DB

func GetClientByType(dbType string) *gorm.DB {
	instance, ok := instances[dbType]
	if !ok {
		log.GetLog().Error(fmt.Sprintf("DB GetClient,no %v instance", dbType))
		panic(fmt.Sprintf("DB GetClient,no %v instance", dbType))
		return nil
	}
	return instance
}

// 获取默认连接 mysql连接
func GetClient() *gorm.DB {
	instance, ok := instances[Mysql]
	if !ok {
		log.GetLog().Error("DB GetClient,no mysql instance")
		panic("DB GetClient,no mysql instance")
		return nil
	}
	return instance
}
