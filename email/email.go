package email

import (
	"github.com/provider-go/pkg/email/qq"
	"github.com/provider-go/pkg/email/typeemail"
)

type SendEmail interface {
	SMTP(fields ...string) error
}

func NewSendEmail(provider string, cfg typeemail.ConfigSendEmail) SendEmail {
	switch provider {
	case "qq":
		return qq.NewEmailQQ(cfg)
	default:

		return qq.NewEmailQQ(cfg)
	}
}
