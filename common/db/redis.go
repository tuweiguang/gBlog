package db

import (
	"context"
	"fmt"
	"gBlog/common/config"
	"gBlog/common/log"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

const redisTimeout = 2 * time.Second

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

	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	err := rdb.Ping(ctx).Err()
	cancel()
	if err != nil {
		rdb.Close()
		log.GetLog().Error(fmt.Sprintf("redis init, ping fail,err = %v", err))
		panic(fmt.Sprintf("redis init, ping fail,err = %v", err))
	}
	return &GRedis{rdb: rdb}
}

type GRedis struct {
	rdb *redis.Client
}

func (g *GRedis) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()
	return g.rdb.Set(ctx, key, value, expiration)
}

func (g *GRedis) Get(key string) *redis.StringCmd {
	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()
	return g.rdb.Get(ctx, key)
}

func (g *GRedis) Incr(key string) *redis.IntCmd {
	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()
	return g.rdb.Incr(ctx, key)
}

func (g *GRedis) HSet(key string, values ...interface{}) *redis.IntCmd {
	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()
	return g.rdb.HSet(ctx, key, values...)
}

func (g *GRedis) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()
	return g.rdb.HIncrBy(ctx, key, field, incr)
}

func (g *GRedis) HGet(key, field string) *redis.StringCmd {
	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()
	return g.rdb.HGet(ctx, key, field)
}

func (g *GRedis) HLen(key string) *redis.IntCmd {
	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()
	return g.rdb.HLen(ctx, key)
}

func (g *GRedis) HGetAll(key string) *redis.StringStringMapCmd {
	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()
	return g.rdb.HGetAll(ctx, key)
}

func (g *GRedis) Close() error {
	return g.rdb.Close()
}
