package auth

import (
	"fmt"
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
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	errs := requests.ValidateRequestIsPhone(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExit(request.Phone),
	})
}

func (sc *SignupController) IsEmailExit(c *gin.Context) {
	type EmailExitRequest struct {
		Email string `json:"email"`
	}
	request := EmailExitRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExit(request.Email),
	})
}
