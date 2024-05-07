package memorystack

import (
	"github.com/provider-go/pkg/stack/typestack"
	"sync"
)

type MemoryStack struct {
	pool *Pool
	lock sync.Mutex
}

// NewMemoryStack 实例化队列，后进先出
func NewMemoryStack(cfg typestack.ConfigStack) (*MemoryStack, error) {
	return &MemoryStack{
		lock: sync.Mutex{},
		pool: new(Pool),
	}, nil
}

// Push 压入队列
func (s *MemoryStack) Push(x interface{}) {
	s.lock.Lock()
	s.pool.Push(x)
	s.lock.Unlock()
}

// Pop 压出队列
func (s *MemoryStack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.pool.Len() > 0 {
		return s.pool.Pop()
	}
	return nil
}

// Len 获取队列长度
func (s *MemoryStack) Len() int {
	return s.pool.Len()
}

type Pool []interface{}

func (p Pool) Len() int { return len(p) }

func (p Pool) Less(i, j int) bool { return true }

func (p Pool) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *Pool) Push(x interface{}) { *p = append(*p, x) }

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
