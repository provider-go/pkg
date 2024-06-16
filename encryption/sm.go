package encryption

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/provider-go/pkg/encryption/sm2"
	"github.com/provider-go/pkg/encryption/sm3"
)

// SM2Sign 私钥签名
func SM2Sign(prik, msg string) string {
	// Private key string conversion type
	dBytes, err := hex.DecodeString(prik)
	if err != nil {
		return ""
	}
	pk := sm2.NewPrivateKey(dBytes)

	signByte, err := pk.Sign(rand.Reader, []byte(msg), nil)
	if err != nil {
		return ""
	} else {
		return hex.EncodeToString(signByte)
	}
}

// SM2Verify 公钥验签
func SM2Verify(pubk, m, s string) bool {
	pubBytes, err := hex.DecodeString(pubk)
	if err != nil {
		return false
	}
	signBytes, err := hex.DecodeString(s)
	if err != nil {
		return false
	}
	pubv := sm2.Decompress(pubBytes)
	msg := []byte(m)
	ok := pubv.Verify(msg, signBytes)
	return ok
}

// SM3Hash SM3
func SM3Hash(msg string) string {
	e := sm3.New()
	/*bf,_ := hex.DecodeString(msg)
	e.Write(bf)*/
	e.Write([]byte(msg))
	return hex.EncodeToString(e.Sum(nil)[:32])
}
