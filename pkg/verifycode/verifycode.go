package verifycode

import (
	"fmt"
	"strings"
	"sync"

	"github.com/VENI-VIDIVICI/plus/pkg/app"
	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/helpers"
	"github.com/VENI-VIDIVICI/plus/pkg/logger"
	"github.com/VENI-VIDIVICI/plus/pkg/mail"
	"github.com/VENI-VIDIVICI/plus/pkg/redis"
	"github.com/VENI-VIDIVICI/plus/pkg/sms"
)

type VerifyCode struct {
	Strore Strore
}

var internalVerfiCode *VerifyCode
var once sync.Once

func NewVerfiCode() *VerifyCode {
	once.Do(func() {
		internalVerfiCode = &VerifyCode{
			Strore: &RedisStroe{
				RedisClient: redis.RedisInstance,
				keyPrefix:   config.GetString("app.name") + ":verifycode",
			},
		}
	})
	return internalVerfiCode
}

func (vc *VerifyCode) SendSMS(phone string) bool {
	// 获取验证码
	code := vc.generateVerifyCode(phone)
	// 方便本地验证
	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	return sms.NewSms().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})
}

func (vc *VerifyCode) CheckAnswer(key, code string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: code})
	if !app.IsProduction() && (strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix")) ||
		strings.HasPrefix(key, config.GetString("verifycode.debug_emial_prefix"))) {
		return true
	}
	return vc.Strore.Verify(key, code, false)
}

func (vc *VerifyCode) SendEmail(email string) error {
	code := vc.generateVerifyCode(email)
	if !app.IsProduction() && strings.HasPrefix(email, config.GetString("verifycode.debug_emial_prefix")) {
		return nil
	}
	content := fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code)
	mail.NewEmail().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{email},
		Subject: "Email 验证码",
		HTML:    []byte(content),
	})
	return nil
}

func (vc *VerifyCode) generateVerifyCode(key string) string {
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))
	if app.IsLocal() {
		code = config.Get("verifycode.debug_code")
	}
	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})
	vc.Strore.Set(key, code)
	return code
}
