package gmsm

import (
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm3"
)

type InstanceGMSM struct {
}

func NewInstanceGMSM() *InstanceGMSM {
	return &InstanceGMSM{}
}

func (c *InstanceGMSM) Hash(data []byte) string {
	e := sm3.New()
	e.Write(data)
	return hex.EncodeToString(e.Sum(nil)[:32])
}
