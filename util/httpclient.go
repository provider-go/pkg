package util

import (
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	Host   string
	Client http.Client
}

func NewHttpClient(host string) *HttpClient {
	return &HttpClient{
		Host: host,
		Client: http.Client{
			Transport: &http.Transport{
				MaxConnsPerHost:     10,
				MaxIdleConnsPerHost: 10,
			},
			Timeout: time.Duration(5) * time.Second,
		},
	}
}

// HttpPostHeaderRequest http请求
func (h *HttpClient) HttpPostHeaderRequest(key, value, params string) (resBody []byte, err error) {
	// 初始化请求
	body := strings.NewReader(params)
	req, err := http.NewRequest("POST", h.Host, body)
	if err != nil {
		return nil, errors.Wrap(err, "Http NewRequest")
	}
	// 执行请求
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(key, value)
	res, err := h.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Client Do")
	}
	defer res.Body.Close()

	// 接收返回结果
	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "io.ReadAll")
	}
	return resBody, nil
}
