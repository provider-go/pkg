package etcd

import (
	"github.com/provider-go/pkg/smcc/typesmcc"
	"testing"
	"time"
)

func Test_KV(t *testing.T) {
	c := typesmcc.ConfigSMCC{
		Endpoints:   []string{"192.168.0.103:2379"},
		DialTimeout: time.Second * 5,
	}
	cli := NewSMCCEtcd(c)
	if cli == nil {
		t.Log("etcd 没连接上")
	}

	err := cli.RegisterService("xm", "qiqi")
	if err != nil {
		t.Log(err)
	}

	for {
		time.Sleep(2 * time.Second)
		service, err := cli.GetService("xm")
		if err != nil {
			t.Log(err)
		}
		t.Log(service)
	}

}
