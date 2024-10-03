package sms

import (
	"sync"

	"github.com/VENI-VIDIVICI/plus/pkg/config"
)

type Message struct {
	Template string
	Data     map[string]string

	Content string
}

type SMS struct {
	Driver Driver
}

var once sync.Once
var internalSMS *SMS

func NewSms() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun_Driver{},
		}
	})
	return internalSMS

}

func (s *SMS) Send(phone string, message Message) bool {
	return s.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
