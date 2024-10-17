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
			authGroup.POST("/signup/email/exist", suc.IsEmailExit)
			vcc := new(auth.VerifyConstroller)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
			authGroup.POST("/signup/phone", suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
		}
	}
}
