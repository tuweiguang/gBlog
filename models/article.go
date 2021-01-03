package models

import (
	"fmt"
	"gBlog/commom/db"
	"gBlog/commom/log"
	"math"
	"time"
)

type Article struct {
	Id         uint `gorm:"primary_key;column:id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	UserId     int64      `gorm:"column:user_id"`
	Title      string     `gorm:"column:title"`
	CategoryId int        `gorm:"column:category_id"`
	Tag        string     `gorm:"column:tag"`
	Content    string     `gorm:"column:content"`
	Status     int        `gorm:"column:status"`
	Pv         int        `gorm:"column:pv"`
	Review     int        `gorm:"column:review"`
	Recommend  int        `gorm:"column:recommend"`
	Like       int        `gorm:"column:like"`
	User       *User
	Category   *Category
}

func (Article) TableName() string {
	return "article"
}

type Paginator struct {
	CurrentPage int //当前页
	PageSize    int //每页数量
	TotalPage   int //总页数
	TotalCount  int //总数量
}

func GenPaginator(page, limit, count int) Paginator {
	var paginator Paginator
	paginator.TotalCount = count
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(limit)))
	paginator.PageSize = limit
	paginator.CurrentPage = page
	return paginator
}

func GetSomeArticle(offset, limit int) []*Article {
	some := make([]*Article, limit)
	err := db.GetMySQL().Raw("SELECT * FROM article as b1 inner join (select id from article limit ?,?) as b2 on b1.id = b2.id", offset, limit).Scan(&some).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetSomeArtcle offset:%v limit:%v error:%v", offset, limit, err))
	}
	return some
}

func CreateArticle(userId int64, title string, categoryId int, tag string, content string) {
	article := Article{
		UserId:     userId,
		Title:      title,
		CategoryId: categoryId,
		Tag:        tag,
		Content:    content,
	}

	result := db.GetMySQL().Create(&article)
	if result.Error != nil {
		log.GetLog().Error(fmt.Sprintf("CreateArticle error:%v", result.Error))
	}
}
