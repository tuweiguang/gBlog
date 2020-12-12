package models

import (
	"math"
	"time"
)

type Article struct {
	Id        int
	Title     string
	Tag       string
	Remark    string
	Desc      string
	Html      string
	Created   time.Time
	Updated   time.Time
	Status    int
	Pv        int
	Review    int
	Recommend int
	Like      int
	User      *User
	Category  *Category
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
