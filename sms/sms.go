package sms

import (
	"github.com/provider-go/pkg/sms/ali"
	"github.com/provider-go/pkg/sms/typesms"
)

type SMS interface {
	Send(fields ...string) error
}

func NewSMS(provider string, cfg typesms.ConfigSMS) SMS {
	switch provider {
	case "ali":
		return ali.NewSMSAli(cfg)
	default:

		return ali.NewSMSAli(cfg)
	}
}
