package validators

import (
	"github.com/VENI-VIDIVICI/plus/pkg/captcha"
	"github.com/VENI-VIDIVICI/plus/pkg/verifycode"
)

func ValidateCaptcha(id string, answer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptch().VerifyCaptcha(id, answer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

func ValidatePasswordConfirm(password string, password_confirm string, errs map[string][]string) map[string][]string {
	if password != password_confirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}

func ValidateVerifyCode(key string, code string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerfiCode().CheckAnswer(key, code); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}
	return errs
}
