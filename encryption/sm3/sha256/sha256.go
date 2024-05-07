package sm3sha256

import (
	"crypto/sha256"
	"fmt"
)

type InstanceSha256 struct {
}

func NewInstanceSha256() *InstanceSha256 {
	return &InstanceSha256{}
}

func (c *InstanceSha256) Hash(data []byte) string {
	h := sha256.Sum256(data)
	return fmt.Sprintf("%x", h)
}
