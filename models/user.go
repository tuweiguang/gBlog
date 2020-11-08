package models

import (
	"gBlog/commom/db"
	"time"
)

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

func GetUser(username string) *User {
	if username == "" {
		return nil
	}

	user := new(User)

	db.GetMySQL().First(user, "name=?", username)
	return user
}
