package leveldbcache

import (
	"github.com/provider-go/pkg/cache/typecache"
	"testing"
)

var cfg = typecache.ConfigCache{
	Addr:     "",
	Password: "",
	DB:       0,
	DBPath:   "../testdata",
}
var ldb, _ = NewCacheLevel(cfg)

func Test_Set(t *testing.T) {
	ldb.Set("qiqi", "KKKKKKKKKK")
}

func Test_Get(t *testing.T) {
	res := ldb.Get("qiqi")
	t.Log(res)
}

func Test_Del(t *testing.T) {
	ldb.Del("qiqi")
}
