package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	"testing"
)

func TestCatIPFS(t *testing.T) {
	sh := shell.NewShell("192.168.0.103:5001")
	hash, err := CatIPFS(sh, "QmTVkPWEgsJKob41Fx9TvdssS4xikYwGwXucPue2qomDMV")
	if err != nil {
		t.Log(err)
	}
	t.Log(string(hash))
}
