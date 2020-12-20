package models

import (
	"fmt"
	"gBlog/commom/db"
	"gBlog/commom/log"
	"time"
)

type User struct {
	Id       int64     `gorm:"primary_key;column:id"`
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

// select * from user as b1 inner join (select id from user limit offset,limit) as b2 on b2.id = b1.id;
func GetSomeUser(offset, limit int) []*User {
	some := make([]*User, limit)
	//db.GetMySQL().Table(getTableName()).Joins(" INNER JOIN  user ON user.id = read_states.target_id ").Where("ats.tenant_id = ? and ats.object_id = ? and ats.object_type = ? AND read_states.status = 'unread'", tenantId, uid, common.EMPLOYEE)
	err := db.GetMySQL().Raw("SELECT * FROM user as b1 inner join (select id from user limit ?,?) as b2 on b1.id = b2.id", offset, limit).Scan(&some).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetSomeUser offset:%v limit:%v error:%v", offset, limit, err))
	}
	return some
}

func GetSomeUserByIds(ids []int64) []*User {
	if ids == nil {
		return nil
	}

	var some []*User
	err := db.GetMySQL().Where("id in (?)", ids).Find(&some).Error
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GetSomeUserByIds error:%v", err))
	}
	return some
}

func IsUserExists(username string) bool {
	if username == "" {
		return false
	}

	user := new(User)

	rows := db.GetMySQL().First(user, "name=?", username).RowsAffected
	if rows != 0 {
		return true
	}

	return false
}

func SaveToDB(user *User) bool {
	if user == nil {
		return false
	}

	err := db.GetMySQL().Save(user).Error
	if err != nil {
		return false
	}
	return true
}
