package redis

import (
	"context"
	"sync"
	"time"

	"github.com/VENI-VIDIVICI/plus/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type RedisWrap struct {
	client redis.Client
	ctx    context.Context
}

var RedisInstance RedisWrap

var once sync.Once

func ConnectRedis(username, address, password string, db int) {
	once.Do(func() {
		NewClient(username, address, password, db)
	})
}

func NewClient(username, address, password string, db int) {
	rds := &RedisWrap{}
	rds.ctx = context.Background()
	rds.client = *redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	err := rds.Ping()
	logger.LogIf(err)
}
func (r RedisWrap) Ping() error {
	_, err := r.client.Ping(r.ctx).Result()
	return err
}

// Decr

func (r RedisWrap) Set(name string, v interface{}, duration int) bool {
	if err := r.client.Set(r.ctx, name, v, time.Duration(duration)).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

func (r RedisWrap) Get(name string) interface{} {
	result, err := r.client.Get(r.ctx, name).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

func (r RedisWrap) Has(name string) bool {
	_, err := r.client.Get(r.ctx, name).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}
func (r RedisWrap) Del(keys ...string) bool {
	if err := r.client.Del(r.ctx, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

func (r RedisWrap) FlushDB() bool {
	if err := r.client.FlushDB(r.ctx).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

// Increment
func (r RedisWrap) DecrAdd(vals ...interface{}) {
	l := len(vals)
	name, ok := vals[0].(string)
	if !ok {
		return
	}
	if l == 1 {
		r.client.Decr(r.ctx, name)
	} else if l == 2 {
		value, ok := vals[1].(int64)
		if !ok {
			return
		}
		r.client.DecrBy(r.ctx, name, value)
	}
}

// Decrement
func (r RedisWrap) IncrD(vals ...interface{}) {
	l := len(vals)
	name, ok := vals[0].(string)
	if !ok {
		return
	}
	if l == 1 {
		r.client.Incr(r.ctx, name)
	} else if l == 2 {
		r.client.IncrBy(r.ctx, name, vals[1].(int64))
	}
}