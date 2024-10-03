package auth

import (
	v1 "github.com/VENI-VIDIVICI/plus/app/http/controllers/api/v1"
	"github.com/VENI-VIDIVICI/plus/app/requests"
	"github.com/VENI-VIDIVICI/plus/pkg/captcha"
	"github.com/VENI-VIDIVICI/plus/pkg/logger"
	"github.com/VENI-VIDIVICI/plus/pkg/response"
	"github.com/VENI-VIDIVICI/plus/pkg/verifycode"
	"github.com/gin-gonic/gin"
)

type VerifyConstroller struct {
	v1.BaseApiController
}

func (ve *VerifyConstroller) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptch().GenerateCaptcha()
	logger.LogIf(err)
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (ve *VerifyConstroller) SendUsingPhone(c *gin.Context) {
	data := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(data, requests.VerifyCodePhone, c); !ok {
		return
	}
	if ok := verifycode.NewVerfiCode().SendSMS(data.Phone); ok {
		response.Success(c)
	} else {
		response.Abort500(c, "发送短信失败")
	}
}

func (ve *VerifyConstroller) SendVerifyCode(c *gin.Context) {

}
