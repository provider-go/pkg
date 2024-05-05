package redis

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
	client := NewCacheRedis(c)
	client.Set("15101131912", "6666")
	t.Log(client.Get("15101131912"))
	client.Del("15101131912")
	t.Log(client.Get("15101131912"))
}
