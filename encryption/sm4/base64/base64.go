package sm4base64

import (
	"encoding/base64"
)

type InstanceBase64 struct {
}

func NewInstanceBase64() *InstanceBase64 {
	return &InstanceBase64{}
}

func (c *InstanceBase64) Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (c *InstanceBase64) Decode(str string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
