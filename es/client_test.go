package es

import (
	"fmt"
	"testing"
	"time"
)

func TestClient_CreateIndex(t *testing.T) {
	client, _ := NewClient("http://60.12.136.237:8002/")
	err := client.CreateIndex("test", `{
	"mappings": {
		"properties": {
			"name": {
				"type": "keyword"
			},
			"uscc": {
				"type": "keyword"
			},
			"address": {
				"type": "keyword"
			}
		}
	},
	"settings": {
		"index": {
			"number_of_shards": 1,
			"number_of_replicas": 1
		}
	}
}`)
	if err != nil {
		t.Logf(err.Error())
	}
}

func TestClient_QueryIndexMappingInfo(t *testing.T) {
	client, _ := NewClient("http://60.12.136.237:8002/")
	indexName := "test"
	result, err := client.QueryIndexMappingInfo(indexName)
	if err != nil {
		t.Logf(err.Error())
	}
	t.Logf("result:%v", result)
}

func TestClient_DeleteIndex(t *testing.T) {
	client, _ := NewClient("http://60.12.136.237:8002/")
	indexName := "test"
	err := client.DeleteIndex(indexName)
	if err != nil {
		t.Logf(err.Error())
	}
}

func TestClient_AddRecord(t *testing.T) {
	client, _ := NewClient("http://60.12.136.237:8002/")
	indexName := "test"
	err := client.AddRecord(indexName, "1", `{"uscc":"J74545dsa484614","name":"微芯研究院","address":"北京市海淀区新中关"}`)
	if err != nil {
		t.Logf(err.Error())
	}
}

func TestClient_BatchAddRecord(t *testing.T) {
	client, _ := NewClient("http://192.168.83.138:9200/")
	indexName := "test"
	body := make([]string, 0)
	body = append(body, `{"name":"xtt","sex":2,"age":30}`)
	body = append(body, `{"name":"xtt","sex":2,"age":40}`)
	n, err := client.BatchAddRecord(indexName, []string{"3", "4"}, body)
	if err != nil {
		t.Fatal("添加失败：", err)
	}
	t.Logf("%v", n)
}

func TestClient_BatchAddRecordTest(t *testing.T) {
	client, _ := NewClient("http://IP:9200/")
	indexName := "test"

	for {
		body := make([]string, 0)
		for i := 0; i < 1000; i++ {
			body = append(body, `{ "name": "xxx", "sex": true, "age": 29}`)
		}
		starttime := time.Now().UnixNano() / 1e6
		n, err := client.BatchAddRecord(indexName, nil, body)
		if err != nil {
			t.Logf(err.Error())
		}
		fmt.Println("耗时：", (time.Now().UnixNano()/1e6)-starttime, "增加交易：", n)
	}
}

func TestClient_UpdateRecord(t *testing.T) {
	client, _ := NewClient("http://IP:9200/")
	indexName := "test"
	n, err := client.UpdateRecord(indexName, `{  "query": {     "match": {      "_id": "uX2aR3QBKSuPjy8yQPSJ"    }  },  "script": {    "source": "ctx._source.age = 36"  }}`)
	if err != nil {
		t.Logf(err.Error())
	}
	t.Logf("%v", n)
}

func TestClient_QueryRecord(t *testing.T) {
	client, _ := NewClient("http://60.12.136.237:8002/")
	indexName := "test"
	retQueryRecord, err := client.QueryRecord(indexName, `{  "query": {    "match_all": {}  },  "from": 0,  "size": 10}`)
	if err != nil {
		t.Logf(err.Error())
	}
	t.Logf("%v", retQueryRecord.String())
}

func TestClient_QueryRecordById(t *testing.T) {
	client, _ := NewClient("http://192.168.83.138:9200/")
	indexName := "test"

	type Account struct {
		Account string `json:"account"`
	}

	result := new(Account)
	err := client.QueryRecordById(indexName, "8iybNnYBmmMJCDkELUma", result)
	if err != nil {
		t.Logf(err.Error())
	}
	t.Logf("%v", result.Account)
}

func TestClient_DeleteRecord(t *testing.T) {
	client, _ := NewClient("http://IP:9200/")
	indexName := "test"
	id := "333CTnQBKSuPjy8yifQ8"
	err := client.DeleteRecord(indexName, id)
	if err != nil {
		t.Logf(err.Error())
	}
}

func TestClient_Count(t *testing.T) {
	c, _ := NewClient("http://192.168.83.138:9200/")
	t.Log(c.Count("test"))
}

func TestClient_BatchDeleteRecord(t *testing.T) {
	arrID := []string{
		"1",
		"2",
	}
	client, _ := NewClient("http://192.168.83.138:9200/")
	indexName := "test"
	err := client.BatchDeleteRecord(indexName, arrID)
	t.Log(err)
}
