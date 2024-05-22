package rpc

import (
	"github.com/provider-go/pkg/rpc/local"
	"github.com/provider-go/pkg/rpc/typerpc"
	"net/http"
)

type LocalRPC interface {
	CallAndReply(http.Handler, *any) error
}

func NewLocalRPC(provider string, cfg typerpc.ConfigRPC) LocalRPC {
	switch provider {
	case "ali":
		return local.NewLocal(cfg)
	default:

		return local.NewLocal(cfg)
	}
}
