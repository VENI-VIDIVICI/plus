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

func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(&request, requests.SignupUsingPhone, c); !ok {
		return
	}
	//
	user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	err := user.Create()
	if err != nil {
		response.Error(c, err)
		return
	}
	response.CreatedJSON(c, gin.H{
		"data": user,
	})
}

func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(&request, requests.SignupEmailExist, c); !ok {
		return
	}

	user := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	err := user.Create()
	if err != nil {
		response.Error(c, err)
		return
	}
	response.CreatedJSON(c, gin.H{
		"data": user,
	})
}
