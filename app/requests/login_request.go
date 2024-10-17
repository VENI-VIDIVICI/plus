package requests

import (
	"github.com/VENI-VIDIVICI/plus/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func LoginByphone(data interface{}, c *gin.Context) map[string][]string {
	// digits:6"
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{"required: Phone 为必填项", "digits: Phone 必须为11位"},
		"verify_code": []string{"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字"},
	}
	errs := validate(data, rules, messages)
	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}
