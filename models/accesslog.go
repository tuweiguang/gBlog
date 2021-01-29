package models

import (
	"fmt"
	"gBlog/commom/db"
	"gBlog/commom/log"
	"time"
)

type AccessLog struct {
	Id     int64     `gorm:"primary_key;column:id"`
	IP     string    `gorm:"column:ip"`
	City   string    `gorm:"column:city"`
	Uri    string    `gorm:"column:uri"`
	Create time.Time `gorm:"column:create"`
}

func (AccessLog) TableName() string {
	return "access_log"
}

func CreateAccessLog(ip, city, uri string) {
	accessLog := AccessLog{
		IP:     ip,
		City:   city,
		Uri:    uri,
		Create: time.Now(),
	}

	result := db.GetMySQL().Create(&accessLog)
	if result.Error != nil {
		log.GetLog().Error(fmt.Sprintf("CreateAccessLog error:%v", result.Error))
	}
}
