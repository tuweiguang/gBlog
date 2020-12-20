package models

import (
	"fmt"
	"gBlog/commom/db"
	"gBlog/commom/log"
)

type Category struct {
	Id     int
	Name   string
	Pid    int
	Sort   int
	Status int
}

func (Category) TableName() string {
	return "category"
}

func GetCategoryById(ids []int) []*Category {
	if ids == nil {
		return nil
	}

	var some []*Category
	err := db.GetMySQL().Where("id in (?)", ids).Find(&some).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetCategoryById error:%v", err))
	}
	return some
}

func GetAllCategory() []*Category {
	var all []*Category
	err := db.GetMySQL().Find(&all).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetAllCategory error:%v", err))
	}
	return all
}
