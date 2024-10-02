package captcha

import (
	"sync"

	"github.com/VENI-VIDIVICI/plus/pkg/app"
	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/redis"
	"github.com/mojocn/base64Captcha"
)

var once sync.Once

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var internalCaptcha *Captcha

func NewCaptch() {
	once.Do(func() {
		internalCaptcha = InitCaptch()
	})
}

func InitCaptch() *Captcha {
	captcha := &Captcha{}
	store := RedisStroe{
		stroeClient: redis.RedisInstance,
		prefix:      config.GetString("app.name") + ":captcha",
	}
	// height int, width int, length int, maxSkew float64, dotCount int
	driver := base64Captcha.NewDriverDigit(config.GetInt("captcha.height"), config.GetInt("captcha.width"),
		config.GetInt("captcha.length"), config.GetFloat64("captcha.maxSkew"), config.GetInt("captcha.dotCount"))
	captcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	return captcha
}

func (c *Captcha) GenerateCaptcha() {
	c.Base64Captcha.Generate()
}

func (c *Captcha) VerifyCaptcha(id string, answer string, clear bool) bool {
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}
	return c.Base64Captcha.Verify(id, answer, clear)
}
