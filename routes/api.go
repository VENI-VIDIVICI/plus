package routes

import (
	"github.com/VENI-VIDIVICI/plus/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(g *gin.Engine) {
	v1 := g.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExit)
			authGroup.POST("/signup/phone/exist", suc.IsEmailExit)
		}
	}
}
