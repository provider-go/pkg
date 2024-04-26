package etcd

import (
	"testing"
	"time"
)

func Test_KV(t *testing.T) {
	cli, err := NewClient([]string{"192.168.0.103:2379"})
	if err != nil {
		t.Log(err)
	}
	err = cli.RegisterService("xm", "qiqi")
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
