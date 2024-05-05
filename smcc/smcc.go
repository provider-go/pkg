package smcc

import (
	"github.com/provider-go/pkg/smcc/etcd"
	"github.com/provider-go/pkg/smcc/typesmcc"
)

type SMCC interface {
	RegisterService(serviceName, serviceAddr string) error
	GetService(key string) ([]string, error)
}

func NewSMCC(provider string, cfg typesmcc.ConfigSMCC) SMCC {
	switch provider {
	case "etcd":
		return etcd.NewSMCCEtcd(cfg)
	default:

		return etcd.NewSMCCEtcd(cfg)
	}
}
