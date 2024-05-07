package sm4

import (
	"github.com/provider-go/pkg/encryption/sm4/base64"
	gmsm4 "github.com/provider-go/pkg/encryption/sm4/gmsm"
)

type SMFour interface {
	Encode(str string) string
	Decode(str string) (string, error)
}

func NewSMFour(provider string) SMFour {
	switch provider {
	case "base64":
		return sm4base64.NewInstanceBase64()
	case "gmsm":
		return gmsm4.NewInstanceGMSM4()
	default:

		return sm4base64.NewInstanceBase64()
	}
}
