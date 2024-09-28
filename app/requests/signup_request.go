package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
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
