package mysql

import (
	"github.com/provider-go/pkg/mysql/typemysql"
	"testing"
)

func TestName(t *testing.T) {
	c := typemysql.ConfigMysql{
		Dsn:          "root:123456@tcp(192.168.0.103:13306)/pyrethrum?parseTime=true&charset=utf8&loc=Local",
		MaxIdleConns: 20,
		MaxOpenConns: 20,
	}
	mysql, err := NewMysql(c)
	if err != nil {
		t.Log(err)
	}
	t.Log(mysql.Name())
}
