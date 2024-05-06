package cache

import (
	"github.com/provider-go/pkg/cache/typecache"
	"testing"
)

func TestName(t *testing.T) {
	c := typecache.ConfigCache{
		Addr:     "192.168.0.103:16379",
		Password: "123456",
		DB:       0,
	}

	cache, _ := NewCache("redis", c)
	cache.Set("xm", "biwow")
	key := cache.Get("xm")
	t.Log(key)
}
