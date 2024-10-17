package auth

import (
	v1 "github.com/VENI-VIDIVICI/plus/app/http/controllers/api/v1"
	"github.com/VENI-VIDIVICI/plus/app/requests"
	"github.com/VENI-VIDIVICI/plus/pkg/auth"
	"github.com/VENI-VIDIVICI/plus/pkg/jwt2"
	"github.com/VENI-VIDIVICI/plus/pkg/response"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseApiController
}

func (lg *LoginController) LoginByPhone(c *gin.Context) {
	// phone string, code string
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(request, requests.LoginByphone, c); !ok {
		return
	}
	user, err := auth.LoginPhone(request.Phone)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	token := jwt2.NewJWT().IssueToken(user.GetStringID(), user.Name)
	response.JSON(c, gin.H{
		"token": token,
	})
}

func (lg *LoginController) LoginByPassword(c *gin.Context) {
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(request, requests.LoginByPassward, c); !ok {
		return
	}
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	token := jwt2.NewJWT().IssueToken(user.GetStringID(), user.Name)
	response.JSON(c, gin.H{
		"token": token,
	})
}
