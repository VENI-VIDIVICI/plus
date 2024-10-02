package captcha

import (
	"errors"
	"fmt"
	"time"

	"github.com/VENI-VIDIVICI/plus/pkg/app"
	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/logger"
	"github.com/VENI-VIDIVICI/plus/pkg/redis"
)

type RedisStroe struct {
	stroeClient *redis.RedisWrap
	prefix      string
}

func (r *RedisStroe) Set(name string, val string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}
	ok := r.stroeClient.Set(fmt.Sprintf("%v-%v", config.GetString("captcha.prefix"), name), val, ExpireTime)
	if !ok {
		logger.ErrorString("captchaStore", "Set", "设置验证码发生错误")
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

func (r *RedisStroe) Get(id string, clear bool) string {
	val := r.stroeClient.Get(fmt.Sprintf("%v_%v", config.GetString("captcha.prefix"), id))
	if clear {
		r.stroeClient.Del(fmt.Sprintf("%v_%v", config.GetString("captcha.prefix")))
	}
	return val.(string)
}

func (r *RedisStroe) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
