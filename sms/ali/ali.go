package ali

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/provider-go/pkg/sms/typesms"
)

type SMSAli struct {
	CFG typesms.ConfigSMS
}

func NewSMSAli(cfg typesms.ConfigSMS) *SMSAli {
	return &SMSAli{
		CFG: cfg,
	}
}

func (a *SMSAli) Send(fields ...string) error {
	config := &openapi.Config{}
	smsClient, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		return err
	}
	request := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(fields[0]),
		SignName:      tea.String(a.CFG.SignName),
		TemplateCode:  tea.String(a.CFG.TemplateCode),
		TemplateParam: tea.String(fields[1]),
	}
	_, err = smsClient.SendSms(request)

	return err
}
