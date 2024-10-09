package mail

import (
	"fmt"
	"net/smtp"

	"github.com/VENI-VIDIVICI/plus/pkg/logger"
	"github.com/jordan-wright/email"
)

type Smtp struct {
}

func (s *Smtp) Send(em Email, config map[string]string) bool {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%v <%v>", em.From.Name, em.From.Address)
	e.To = em.To
	e.Bcc = em.Bcc
	e.Cc = em.Cc
	e.Subject = em.Subject
	e.Text = em.Text
	e.HTML = em.HTML
	logger.DebugJSON("发送邮件", "发送详情", e)
	err := e.Send(fmt.Sprintf("%v:%v", config["host"], config["port"]),
		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)
	if err != nil {
		logger.ErrorString("发送邮件", "发送出错", err.Error())
		return false
	}
	logger.ErrorString("发送邮件", "发送成功", "")
	return true
}
