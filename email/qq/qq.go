package qq

import (
	"crypto/tls"
	"github.com/provider-go/pkg/email/typeemail"
	"gopkg.in/gomail.v2"
)

type EmailQQ struct {
	CFG    typeemail.ConfigSendEmail
	client *gomail.Dialer
}

func NewEmailQQ(cfg typeemail.ConfigSendEmail) *EmailQQ {
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.User, cfg.AuthCode)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &EmailQQ{
		CFG:    cfg,
		client: d,
	}
}

// SMTP fields[0] 收信人 fields[1] 邮件标题 fields[2] 邮件内容
func (a *EmailQQ) SMTP(fields ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", a.CFG.User)
	m.SetHeader("To", fields[0])
	m.SetHeader("Subject", fields[1])
	m.SetBody("text/html", fields[2])

	err := a.client.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
