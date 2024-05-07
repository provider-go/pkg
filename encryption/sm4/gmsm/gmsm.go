package gmsm4

import (
	"encoding/hex"
	"github.com/provider-go/pkg/util"
)

type InstanceGMSM4 struct {
}

func NewInstanceGMSM4() *InstanceGMSM4 {
	return &InstanceGMSM4{}
}

func (g *InstanceGMSM4) Encode(str string) string {
	if len(str) < 16 {
		num := 16 - len(str)
		str = util.RightPadZeros(str, num)
	}
	key := []byte("1234567890abcdef")
	c, err := NewCipher(key)
	if err != nil {
		return ""
	}
	dst := make([]byte, 16)
	c.Encrypt(dst, []byte(str))
	return hex.EncodeToString(dst)
}

func (g *InstanceGMSM4) Decode(str string) (string, error) {
	b, err := hex.DecodeString(str)
	if err != nil {
		return "", err
	}
	key := []byte("1234567890abcdef")
	c, err := NewCipher(key)
	if err != nil {
		return "", err
	}

	dst := make([]byte, 16)
	c.Decrypt(dst, b)
	return string(dst), nil
}
