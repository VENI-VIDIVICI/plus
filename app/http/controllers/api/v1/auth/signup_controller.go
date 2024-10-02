package auth

import (
	v1 "github.com/VENI-VIDIVICI/plus/app/http/controllers/api/v1"
	"github.com/VENI-VIDIVICI/plus/app/models/user"
	"github.com/VENI-VIDIVICI/plus/app/requests"
	"github.com/VENI-VIDIVICI/plus/pkg/response"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExit(c *gin.Context) {
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(&request, requests.SignupPhoneExit, c); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExit(request.Phone),
	})
}

func (sc *SignupController) IsEmailExit(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(&request, requests.SignupEmailExist, c); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": user.IsEmailExit(request.Email),
	})
}
