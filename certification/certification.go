package certification

import (
	"github.com/provider-go/pkg/certification/ali"
	"github.com/provider-go/pkg/certification/typecert"
)

type Certification interface {
	Send(fields ...string) string
}

func NewCertification(provider string, cfg typecert.ConfigCertification) Certification {
	switch provider {
	case "ali":
		return ali.NewCertificationAli(cfg)
	default:

		return ali.NewCertificationAli(cfg)
	}
}
