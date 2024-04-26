package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type HttpConPool struct {
	Conn *http.Client
}

var once sync.Once
var hpool *HttpConPool

// NewHttpPool 创建连接池
func NewHttpPool(maxConnPer, maxIdleConnPer int, duration int64) *HttpConPool {
	once.Do(func() {
		hpool = new(HttpConPool)
		hpool.Conn = &http.Client{
			Transport: &http.Transport{
				MaxConnsPerHost:     maxConnPer,
				MaxIdleConnsPerHost: maxIdleConnPer,
			},
			Timeout: time.Duration(duration) * time.Second,
		}
	})
	return hpool
}

// Request send a http request of post or get
func (h *HttpConPool) Request(url string, method string, data string, header map[string]string) (string, error, int) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err, 0
	}

	for h, v := range header {
		req.Header.Set(h, v)
	}

	// 如果es有密码则增加
	if viper.GetString("es.user") != "" && viper.GetString("es.pass") != "" {
		req.SetBasicAuth(viper.GetString("es.user"), viper.GetString("es.pass"))
	}
	response, err := h.Conn.Do(req)

	if err != nil {
		return "", err, 0
	} else if response != nil {
		defer response.Body.Close()

		r_body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err, response.StatusCode
		} else {
			return string(r_body), nil, response.StatusCode
		}
	} else {
		return "", nil, 0
	}
}

const ConnTimeOut = 3
const OKCode = 200

// HTTPRequest http请求
func HTTPRequest(method, url, body string, OKCode int, result interface{}) error {
	hp := NewHttpPool(0, 0, ConnTimeOut)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	res, err, statusCode := hp.Request(url, method, body, header)
	if err != nil {
		return fmt.Errorf("HTTPRequest http.NewRequest err:%v", err.Error())
	}

	if statusCode != OKCode && statusCode != 200 {
		errResult := new(ResponseError)
		err = json.Unmarshal([]byte(res), errResult)
		if err != nil {
			return fmt.Errorf("HTTPRequest response data error json.Unmarshal:%v", err.Error())
		}
		return fmt.Errorf("HTTPRequest request res.StatusCode != %v, res.StatusCode=%v, errType:%v, errReason:%v", OKCode, statusCode, errResult.Error.Type, errResult.Error.Reason)
	}
	err = json.Unmarshal([]byte(res), result)
	if err != nil {
		return fmt.Errorf("HTTPRequest json.Unmarshal err:%v", err.Error())
	}
	return nil
}
