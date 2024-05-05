package redis

import (
	"github.com/go-redis/redis"
	"github.com/provider-go/pkg/cache/typecache"
	"github.com/provider-go/pkg/logger"
)

type CacheRedis struct {
	db *redis.Client // 数据库句柄
}

func NewCacheRedis(cfg typecache.ConfigCache) *CacheRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		logger.Error("redis:", "step", "NewCacheRedis", "err", err)
		return nil
	}
	return &CacheRedis{db: client}
}

// Set 写入数据
func (r *CacheRedis) Set(key, value string) {
	r.db.Set(key, value, 0)
}

// Get 读数据
func (r *CacheRedis) Get(key string) string {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Redis Get", "step", "defer", "err", err)
		}
	}()
	value, err := r.db.Get(key).Result()
	if err != nil || value == "" {
		return ""
	}
	return value
}

// Del 删除数据
func (r *CacheRedis) Del(key string) {
	r.db.Del(key)
}
