package mail

import (
	"sync"

	"github.com/VENI-VIDIVICI/plus/pkg/config"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Cc      []string
	Subject string
	Bcc     []string
	Text    []byte
	HTML    []byte
}

type Mailer struct {
	Driver Driver
}

var innerEmail *Mailer
var once sync.Once

func NewEmail() *Mailer {
	once.Do(func() {
		innerEmail = &Mailer{
			Driver: &Smtp{},
		}
	})
	return innerEmail
}

func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, config.GetStringMapString("mail.smtp"))
}
