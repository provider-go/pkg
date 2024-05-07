package stack

import (
	memorystack "github.com/provider-go/pkg/stack/memory"
	"github.com/provider-go/pkg/stack/typestack"
)

type Stack interface {
	Push(x interface{})
	Pop() interface{}
	Len() int
}

func NewStack(provider string, cfg typestack.ConfigStack) (Stack, error) {
	switch provider {
	case "memory":
		return memorystack.NewMemoryStack(cfg)
	default:

		return memorystack.NewMemoryStack(cfg)
	}
}
