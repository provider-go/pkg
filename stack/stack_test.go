package stack

import (
	"github.com/provider-go/pkg/stack/typestack"
	"testing"
)

func TestName(t *testing.T) {
	c := typestack.ConfigStack{
		Addr:     "192.168.0.103:16379",
		Password: "123456",
		DB:       0,
	}

	sk, _ := NewStack("memory", c)
	sk.Push(1111)
	sk.Push(2222)
	sk.Push(3333)
	sk.Push(4444)
	t.Log(sk.Len())
	t.Log(sk.Pop())
	t.Log(sk.Len())

}
