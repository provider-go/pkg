package sm3md5

import (
	"crypto/md5"
	"encoding/hex"
)

type InstanceMd5 struct {
}

func NewInstanceMd5() *InstanceMd5 {
	return &InstanceMd5{}
}

func (c *InstanceMd5) Hash(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
