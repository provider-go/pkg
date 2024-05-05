package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	"strings"
	"testing"
)

func TestUploadIPFS(t *testing.T) {
	sh := shell.NewShell("192.168.0.103:5001")
	hash, err := UploadIPFS(sh, strings.NewReader("hello world"))
	if err != nil {
		t.Log(err)
	}
	t.Log(hash)
}
