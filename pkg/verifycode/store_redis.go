package verifycode

import (
	"time"

	"github.com/VENI-VIDIVICI/plus/pkg/app"
	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/redis"
)

type RedisStroe struct {
	RedisClient *redis.RedisWrap
	keyPrefix   string
}

// var once sync.Once

// var innerStore *RedisStroe

// func NewStore() *RedisStroe {
// 	once.Do(func() {
// 		innerStore = &RedisStroe{
// 			RedisClient: redis.RedisInstance,
// 			keyPrefix:   config.GetString("app.name") + ":verifycode",
// 		}
// 	})
// 	return innerStore
// }

func (s *RedisStroe) Set(key, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt("verifycode.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt("verifycode.debug_expire_time"))
	}
	return s.RedisClient.Set(s.keyPrefix+key, value, ExpireTime)
}

func (s *RedisStroe) Get(key string, clear bool) (value string) {
	// return s.RedisClient.Set()
	key = s.keyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val.(string)
}

func (s *RedisStroe) Verify(id, answer string, clear bool) bool {
	key := s.keyPrefix + id
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val.(string) == answer
}
