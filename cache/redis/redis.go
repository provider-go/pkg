package rediscache

import (
	"github.com/go-redis/redis"
	"github.com/provider-go/pkg/cache/typecache"
	"github.com/provider-go/pkg/logger"
	"time"
)

type CacheRedis struct {
	db *redis.Client // 数据库句柄
}

func NewCacheRedis(cfg typecache.ConfigCache) (*CacheRedis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &CacheRedis{db: client}, nil
}

// Set 写入数据
func (r *CacheRedis) Set(key, value string, expiration time.Duration) {
	r.db.Set(key, value, time.Second*expiration)
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
