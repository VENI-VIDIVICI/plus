package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func SignupPhoneExit(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	message := govalidator.MapData{
		"phone": []string{"required:请输入手机号码", "digits:手机号码格式不正确"},
	}
	return validate(&data, rules, message)
}

func SignupEmailExist(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data, rules, messages)
}
