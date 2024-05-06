package cache

import (
	"github.com/provider-go/pkg/cache/redis"
	"github.com/provider-go/pkg/cache/typecache"
)

type Cache interface {
	Set(key, value string)
	Get(key string) string
	Del(key string)
}

func NewCache(provider string, cfg typecache.ConfigCache) (Cache, error) {
	switch provider {
	case "redis":
		return redis.NewCacheRedis(cfg)
	default:

		return redis.NewCacheRedis(cfg)
	}
}
