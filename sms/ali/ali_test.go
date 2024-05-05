package ali

import (
	"github.com/provider-go/pkg/sms/typesms"
	"testing"
)

func TestSend(t *testing.T) {
	c := typesms.ConfigSMS{
		AccessKeyId:     "",
		AccessKeySecret: "",
		Endpoint:        "dysmsapi.aliyuncs.com",
		SignName:        "XXX",
		TemplateCode:    "SMS_245100000",
	}
	aliSMS := NewSMSAli(c)
	err := aliSMS.Send("15101131912", "{ \"code\" : 123456}")
	if err != nil {
		t.Log(err)
	}
}
