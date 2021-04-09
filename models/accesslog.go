package models

import (
	"fmt"
	"gBlog/common/db"
	"gBlog/common/log"
	"time"
)

type AccessLog struct {
	Id      int64     `gorm:"primary_key;column:id"`
	IP      string    `gorm:"column:ip"`
	Country string    `grom:"country"`
	City    string    `gorm:"column:city"`
	ISP     string    `gorm:"column:ISP"`
	Uri     string    `gorm:"column:uri"`
	Create  time.Time `gorm:"column:create"`
}

func (AccessLog) TableName() string {
	return "access_log"
}

func CreateAccessLog(ip, country, city, isp, uri string) {
	accessLog := AccessLog{
		IP:      ip,
		Country: country,
		City:    city,
		ISP:     isp,
		Uri:     uri,
		Create:  time.Now(),
	}

	result := db.GetMySQL().Create(&accessLog)
	if result.Error != nil {
		log.GetLog().Error(fmt.Sprintf("CreateAccessLog error:%v", result.Error))
	}
}
