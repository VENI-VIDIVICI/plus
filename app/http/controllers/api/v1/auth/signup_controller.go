package auth

import (
	"net/http"

	v1 "github.com/VENI-VIDIVICI/plus/app/http/controllers/api/v1"
	"github.com/VENI-VIDIVICI/plus/app/models/user"
	"github.com/VENI-VIDIVICI/plus/app/requests"
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

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExit(request.Phone),
	})
}

func (sc *SignupController) IsEmailExit(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(&request, requests.SignupEmailExist, c); !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExit(request.Email),
	})
}
