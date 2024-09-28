package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// type ValidateHandle func
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(data interface{}, handle ValidatorFunc, c *gin.Context) bool {
	if err := c.ShouldBindJSON(data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return false
	}
	errs := handle(data, c)
	if len(errs) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opt := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	return govalidator.New(opt).ValidateStruct()
}
