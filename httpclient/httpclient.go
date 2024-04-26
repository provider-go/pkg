package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type HttpClient struct {
	Host   string
	Client http.Client
}

var (
	once   sync.Once
	client HttpClient
)

func Client(host string) HttpClient {
	once.Do(func() {
		client = HttpClient{
			Host: host,
			Client: http.Client{
				Transport: &http.Transport{
					MaxConnsPerHost:     50,
					MaxIdleConnsPerHost: 50,
				},
				Timeout: time.Duration(5) * time.Second,
			},
		}
	})
	return client
}

// Request send a http request of post or get
func (h HttpClient) Request(method, path, data string, header map[string]string) (string, error) {
	req, err := http.NewRequest(method, StrConcat(h.Host, "/", path), bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err
	}

	for h, v := range header {
		req.Header.Set(h, v)
	}

	response, err := h.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func StrConcat(args ...string) string {
	var build strings.Builder
	for _, v := range args {
		build.WriteString(v)
	}
	return build.String()
}
