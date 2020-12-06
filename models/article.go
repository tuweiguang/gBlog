package models

import "time"

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
