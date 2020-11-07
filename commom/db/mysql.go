package db

import (
	"fmt"
	"gBlog/commom/config"
	"gBlog/commom/log"

	"github.com/jinzhu/gorm"
)

type MySQL struct{}

func (m *MySQL) init(dbI SQL) *gorm.DB {
	c, ok := dbI.(config.DB)
	if !ok {
		log.GetLog().Error("mysql init fail")
		panic("mysql init fail")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DbName, c.DbCharset)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Log.Error(err.Error())
		panic("mysql init open fail")
	}
	return db
}
