package rpc

import (
	"github.com/provider-go/pkg/rpc/tests"
	"github.com/provider-go/pkg/rpc/typerpc"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	c := typerpc.ConfigRPC{
		Method: http.MethodGet,
		URL:    "/ping?id=666",
		Data:   "",
		Header: make(map[string]string),
	}

	rpc := NewLocalRPC("local", c)

	var res any
	err := rpc.CallAndReply(tests.NewApp().R, &res)
	if err != nil {
		t.Log(err)
	} else {
		s := res.(string)
		t.Log(s)
	}

}
