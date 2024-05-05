package file

import (
	"bytes"
	"github.com/provider-go/pkg/file/typefile"
	"testing"
)

func TestName(t *testing.T) {
	c := typefile.ConfigStorageFile{Endpoints: "192.168.0.103:5001"}
	client := NewStorageFile("ipfs", c)
	hash, err := client.Upload(bytes.NewBuffer([]byte("qiqi")))
	if err != nil {
		t.Log(err)
	}
	t.Log(hash)
	res, err := client.Download(hash)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(res))
}
