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

func ValidateRequestIsPhone(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	message := govalidator.MapData{
		"phone": []string{"required:请输入手机号码", "digits:手机号码格式不正确"},
	}
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      message,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	}
	return govalidator.New(opts).ValidateStruct()
}

func ValidateRequestIsEmail(data interface{}, ctx *gin.Context) map[string][]string {
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
	opt := govalidator.Options{
		Data:          data,
		Messages:      messages,
		Rules:         rules,
		TagIdentifier: "valid",
	}
	return govalidator.New(opt).ValidateStruct()
}
