package memorycache

import (
	"encoding/json"
	"github.com/provider-go/pkg/cache/typecache"
	"testing"
)

var cfg = typecache.ConfigCache{
	Addr:     "",
	Password: "",
	DB:       0,
	DBPath:   "../testdata",
}
var mdb, _ = NewCacheMemory(cfg)

func Test_Set_Get(t *testing.T) {
	// set
	mdb.Set("qiqi", "KKKKKKKKKK")
	b, _ := json.Marshal(mdb.db)
	t.Log(string(b))
	// get
	res := mdb.Get("qiqi")
	t.Log(res)
	// del
	mdb.Del("qiqi")
	// get
	res = mdb.Get("qiqi")
	t.Log(res)
}
