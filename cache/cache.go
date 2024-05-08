package cache

import (
	leveldbcache "github.com/provider-go/pkg/cache/leveldb"
	memorycache "github.com/provider-go/pkg/cache/memory"
	rediscache "github.com/provider-go/pkg/cache/redis"
	"github.com/provider-go/pkg/cache/typecache"
	"time"
)

type Cache interface {
	Set(key, value string, expiration time.Duration)
	Get(key string) string
	Del(key string)
}

func NewCache(provider string, cfg typecache.ConfigCache) (Cache, error) {
	switch provider {
	case "redis":
		return rediscache.NewCacheRedis(cfg)
	case "level":
		return leveldbcache.NewCacheLevelDB(cfg)
	case "memory":
		return memorycache.NewCacheMemory(cfg)
	default:

		return rediscache.NewCacheRedis(cfg)
	}
}
