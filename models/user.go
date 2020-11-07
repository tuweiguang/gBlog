package models

import "time"

type User struct {
	Id       int       `gorm:"primary_key;column:id"`
	Name     string    `gorm:"column:name"`
	Password string    `gorm:"column:password"`
	Email    string    `gorm:"column:email"`
	Created  time.Time `gorm:"column:created"`
	Status   int       `gorm:"column:status"`
}

func (User) TableName() string {
	return "user"
}
