package memorystack

import (
	"github.com/provider-go/pkg/stack/typestack"
	"testing"
)

var cfg = typestack.ConfigStack{
	Addr:     "",
	Password: "",
	DB:       0,
}
var stack, _ = NewMemoryStack(cfg)

func TestStackPool_Push(t *testing.T) {
	stack.Push("111111")
	stack.Push("222222")
	stack.Push("333333")
	stack.Push("444444")
	t.Log(stack.Len())
	t.Log(stack.Pop())
	t.Log(stack.Len())
}
