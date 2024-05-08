package memorycache

import (
	"github.com/provider-go/pkg/cache/typecache"
	"sync"
	"time"
)

// 结构体
type CacheMemoryDB struct {
	lock sync.Mutex
	db   map[string]string
}

// NewCacheMemory 初始化数据库
func NewCacheMemory(cfg typecache.ConfigCache) (*CacheMemoryDB, error) {
	return &CacheMemoryDB{
		db: make(map[string]string),
	}, nil
}

// Set 写方法
func (db *CacheMemoryDB) Set(key, value string, expiration time.Duration) {
	db.lock.Lock()
	db.db[key] = value
	db.lock.Unlock()
}

// Get 读方法
func (db *CacheMemoryDB) Get(key string) string {
	db.lock.Lock()
	defer db.lock.Unlock()

	if value, ok := db.db[key]; ok {
		return value
	}
	return ""
}

// Del 删除制定键值
func (db *CacheMemoryDB) Del(key string) {
	db.lock.Lock()
	delete(db.db, key)
	db.lock.Unlock()
}
