package db

import (
	"context"
	"gBlog/commom/config"
	"gBlog/commom/log"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type Redis struct{}

func (r *Redis) init(dbI SQL) interface{} {
	c, ok := dbI.(config.DB)
	if !ok {
		log.GetLog().Error("mysql init fail")
		panic("mysql init fail")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + strconv.Itoa(c.Port),
		Password: c.Password, // no password set
		DB:       c.DbNum,    // use default DB
	})

	return &GRedis{rdb: rdb}
}

var ctx = context.Background()

type GRedis struct {
	rdb *redis.Client
}

func (g *GRedis) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return g.rdb.Set(ctx, key, value, expiration)
}

func (g *GRedis) Get(key string) *redis.StringCmd {
	return g.rdb.Get(ctx, key)
}

func (g *GRedis) Close() error {
	return g.rdb.Close()
}
