package session

import (
	"gBlog/commom/config"
	"gBlog/commom/db"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

var rOnce sync.Once
var redisMgr RedisMgr

type RedisMgr struct {
	sessionExpire int // session过期时间
	redis         *db.GRedis
}

func NewRedisMgr() *RedisMgr {
	rOnce.Do(func() {
		redisMgr.redis = db.GetRedis()
		redisMgr.sessionExpire = config.GetSessionConfig().Expire
	})

	return &redisMgr
}

func (r *RedisMgr) CheckSession(sessionId string) int {
	_, err := r.redis.Get(sessionId).Result()
	if err != nil {
		return SessionNoexist
	}

	return SessionExist
}

func (r *RedisMgr) CreateSessoin() string {
	// 构造一个sessionId
	sessionId := uuid.NewV4().String()

	r.redis.Set(sessionId, 1, time.Duration(r.sessionExpire)*time.Second)
	return sessionId
}
