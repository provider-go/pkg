package leveldbcache

import (
	"github.com/provider-go/pkg/cache/typecache"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type CacheLevelDB struct {
	db *leveldb.DB
}

func NewCacheLevelDB(cfg typecache.ConfigCache) (*CacheLevelDB, error) {
	// 打开数据库并定义相关参数
	db, err := leveldb.OpenFile(cfg.DBPath, &opt.Options{
		Compression:         opt.SnappyCompression,
		WriteBuffer:         32 * opt.MiB,
		CompactionTableSize: 2 * opt.MiB,               // 定义数据文件最大存储
		Filter:              filter.NewBloomFilter(10), // bloom过滤器
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(cfg.DBPath, nil)
	}
	if err != nil {
		return nil, err
	}

	// 结构体赋值并返回
	return &CacheLevelDB{db: db}, nil
}

// Set 数据库写操作
func (db *CacheLevelDB) Set(key, value string) {
	_ = db.db.Put([]byte(key), []byte(value), nil)
}

// Get 数据库读操作
func (db *CacheLevelDB) Get(key string) string {
	data, _ := db.db.Get([]byte(key), nil)

	return string(data)
}

// Del 数据库删除操作
func (db *CacheLevelDB) Del(key string) {
	_ = db.db.Delete([]byte(key), nil)
}

// NewIterator 数据库迭代器
func (db *CacheLevelDB) NewIterator() iterator.Iterator {
	return db.db.NewIterator(nil, nil)
}

// Close 关闭数据库
func (db *CacheLevelDB) Close() error {
	if err := db.db.Close(); err != nil {
		return err
	}
	return nil
}
