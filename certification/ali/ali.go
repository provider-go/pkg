package ali

import (
	"github.com/provider-go/pkg/certification/typecert"
	"github.com/provider-go/pkg/logger"
	"github.com/provider-go/pkg/util"
	"net/url"
)

type CertificationAli struct {
	CFG typecert.ConfigCertification
}

func NewCertificationAli(cfg typecert.ConfigCertification) *CertificationAli {
	return &CertificationAli{
		CFG: cfg,
	}
}

// Send fields[0] 身份证 fields[1] 姓名
func (a *CertificationAli) Send(fields ...string) string {
	params := "idcard=" + fields[0] + "&name=" + url.QueryEscape(fields[1])
	client := util.NewHttpClient(a.CFG.Endpoint)
	resBody, err := client.HttpPostHeaderRequest("Authorization", "APPCODE "+a.CFG.Appcode+"", params)
	if err != nil {
		logger.Error("CertificationAli", "step", "HttpPostHeaderRequest", "err", err)
		return ""
	}
	return string(resBody)
}
