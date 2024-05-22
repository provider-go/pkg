package local

import (
	"bytes"
	"errors"
	"github.com/provider-go/pkg/rpc/typerpc"
	"net/http"
	"net/http/httptest"
	"strconv"
)

type Local struct {
	*http.Request
}

func NewLocal(cfg typerpc.ConfigRPC) *Local {
	req, err := http.NewRequest(cfg.Method, cfg.URL, bytes.NewBuffer([]byte(cfg.Data)))
	if err != nil {
		return nil
	}

	for h, v := range cfg.Header {
		req.Header.Set(h, v)
	}
	return &Local{req}
}

func (l *Local) CallAndReply(r http.Handler, reply *any) error {
	res := httptest.NewRecorder()
	r.ServeHTTP(res, l.Request)

	if res.Code != 200 {
		return errors.New("请求状态码为" + strconv.Itoa(res.Code))
	}
	*reply = res.Body.String()
	return nil
}
