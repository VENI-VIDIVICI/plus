package validators

import "github.com/VENI-VIDIVICI/plus/pkg/captcha"

func ValidateCaptcha(id string, answer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptch().VerifyCaptcha(id, answer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}
