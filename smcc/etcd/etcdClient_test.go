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
	cli, err := NewSMCCEtcd(c)
	if err != nil {
		t.Log("etcd 没连接上")
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

func TestConfig(t *testing.T) {
	c := typesmcc.ConfigSMCC{
		Endpoints:   []string{"192.168.0.103:2379"},
		DialTimeout: time.Second * 5,
	}
	cli, err := NewSMCCEtcd(c)
	if err != nil {
		t.Log("etcd 没连接上")
	}

	err = cli.SetConfig("mysql.dsn", "111111")
	if err != nil {
		t.Log(err)
	}
	err = cli.SetConfig("mysql.username", "222222")
	if err != nil {
		t.Log(err)
	}
	err = cli.SetConfig("mysql.pass", "33333")
	if err != nil {
		t.Log(err)
	}
	err = cli.SetConfig("mysql.port", "44444")
	if err != nil {
		t.Log(err)
	}

	res, err := cli.GetConfig("name")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
	res, err = cli.GetConfig("mysql.dsn")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
	res, err = cli.GetConfig("mysql.username")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
	res, err = cli.GetConfig("mysql.pass")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
	res, err = cli.GetConfig("mysql.port")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}
