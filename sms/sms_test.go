package sms

import (
	"github.com/provider-go/pkg/sms/typesms"
	"testing"
)

func TestName(t *testing.T) {
	c := typesms.ConfigSMS{
		AccessKeyId:     "",
		AccessKeySecret: "",
		Endpoint:        "dysmsapi.aliyuncs.com",
		SignName:        "XXX",
		TemplateCode:    "SMS_245100000",
	}
	s := NewSMS("", c)
	err := s.Send("15101131912", "{ \"code\" : 123456}")
	if err != nil {
		t.Log(err)
	}
}
