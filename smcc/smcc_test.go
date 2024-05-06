package smcc

import (
	"github.com/provider-go/pkg/smcc/typesmcc"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	c := typesmcc.ConfigSMCC{
		Endpoints:   []string{"192.168.0.103:2379"},
		DialTimeout: time.Second * 5,
	}
	s, _ := NewSMCC("etcd", c)
	err := s.RegisterService("xm", "qiqi")
	if err != nil {
		t.Log(err)
	}
	res, err := s.GetService("xm")
	if err != nil {
		t.Log(err)
	}
	for _, value := range res {
		t.Log("Service: " + value)
	}
}
