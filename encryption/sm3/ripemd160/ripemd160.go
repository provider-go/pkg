package sm3ripemd160

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

// Ripemd160 和 sha256 双hash

type InstanceRipemd160 struct {
}

func NewInstanceRipemd160() *InstanceRipemd160 {
	return &InstanceRipemd160{}
}

// Hash returns the RIPEMD160 hash of the SHA-256 HASH of the given data.
func (c *InstanceRipemd160) Hash(data []byte) string {
	h := sha256.Sum256(data)
	return fmt.Sprintf("%x", Ripemd160h(h[:]))
}

// Ripemd160h returns the RIPEMD160 hash of the given data.
func Ripemd160h(data []byte) []byte {
	h := ripemd160.New()
	h.Write(data)
	return h.Sum(nil)
}
