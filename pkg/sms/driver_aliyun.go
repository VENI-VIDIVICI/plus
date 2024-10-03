package sms

import (
	"encoding/json"

	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
	"github.com/VENI-VIDIVICI/plus/pkg/logger"
)

type Aliyun_Driver struct{}

var (
	gatewayUrl = "http://dysmsapi.aliyuncs.com/"
	// accessKeyId     = "LTAIbTnPbawglLIQ"
	// accessKeySecret = ""
	// phoneNumbers    = "13544285**2"
	// signName        = "坤Kenmy"
	// templateCode    = "SMS_82045083"
	// templateParam   = "{\"code\":\"1234\"}"
)

func (alyun *Aliyun_Driver) Send(phone string, message Message, config map[string]string) bool {
	smsClient := aliyunsmsclient.New(gatewayUrl)
	templateParams, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("Sms", "Send", err.Error())
		return false
	}
	result, err := smsClient.Execute(config["access_keyId"], config["access_keySecret"], phone, config["sign_name"], message.Template, string(templateParams))
	if err != nil {
		logger.ErrorJSON("Sms", "Execute", config)
		return false
	}
	resultJSON, err := json.Marshal(result)
	logger.DebugJSON("短信[阿里云]", "请求内容", smsClient.Request)
	logger.DebugJSON("短信[阿里云]", "接口响应", result)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析响应 JSON 错误", err.Error())
		return false
	}
	if result.IsSuccessful() {
		logger.DebugString("短信[阿里云]", "发信成功", "")
		return true
	} else {
		logger.ErrorString("短信[阿里云]", "服务商返回错误", string(resultJSON))
		return false
	}
}
