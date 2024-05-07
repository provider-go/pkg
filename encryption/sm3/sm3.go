package sm3

import (
	"github.com/provider-go/pkg/encryption/sm3/gmsm"
	sm3md5 "github.com/provider-go/pkg/encryption/sm3/md5"
	sm3ripemd160 "github.com/provider-go/pkg/encryption/sm3/ripemd160"
	sm3sha256 "github.com/provider-go/pkg/encryption/sm3/sha256"
)

type SMThree interface {
	Hash(data []byte) string
}

func NewSMThree(provider string) SMThree {
	switch provider {
	case "md5":
		return sm3md5.NewInstanceMd5()
	case "sha256":
		return sm3sha256.NewInstanceSha256()
	case "ripemd160":
		return sm3ripemd160.NewInstanceRipemd160()
	case "gmsm":
		return gmsm.NewInstanceGMSM()
	default:

		return gmsm.NewInstanceGMSM()
	}
}
