package db

import (
	"context"
	"fmt"
	"gBlog/commom/config"
	"gBlog/commom/log"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type Redis struct{}

func (r *Redis) init(dbI SQL) interface{} {
	c, ok := dbI.(*config.DB)
	if !ok {
		log.GetLog().Error("redis init, config error")
		panic("redis init, config error")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Host + ":" + strconv.Itoa(c.Port),
		Password: c.Password, // no password set
		DB:       c.DbNum,    // use default DB
	})

	if rdb == nil {
		log.GetLog().Error("redis init fail")
		panic("redis init fail")
	}

	err := rdb.Ping(ctx).Err()
	if err != nil {
		rdb.Close()
		log.GetLog().Error(fmt.Sprintf("redis init, ping fail,err = %v", err))
		panic(fmt.Sprintf("redis init, ping fail,err = %v", err))
	}
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

func (g *GRedis) HSet(key string, values ...interface{}) *redis.IntCmd {
	return g.rdb.HSet(ctx, key, values...)
}

func (g *GRedis) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	return g.rdb.HIncrBy(ctx, key, field, incr)
}

func (g *GRedis) HGet(key, field string) *redis.StringCmd {
	return g.rdb.HGet(ctx, key, field)
}

func (g *GRedis) Close() error {
	return g.rdb.Close()
}
