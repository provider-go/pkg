package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type ClientConn struct {
	client *clientv3.Client
}

// NewClient 创建etcd链接
func NewClient(addr []string) (*ClientConn, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	if client, err := clientv3.New(conf); err == nil {
		return &ClientConn{
			client: client,
		}, nil
	} else {
		return nil, err
	}
}

// RegisterService 服务注册
func (c *ClientConn) RegisterService(serviceName, serviceAddr string) error {
	// 创建一个新的lease
	resp, err := c.client.Grant(context.TODO(), 5) // 5秒后过期
	if err != nil {
		panic(err)
	}
	leaseID := resp.ID
	// 写入服务注册信息
	_, err = c.client.Put(context.Background(), serviceName, serviceAddr, clientv3.WithLease(leaseID))
	if err != nil {
		return err
	}
	// 续租
	ch, kaerr := c.client.KeepAlive(context.TODO(), leaseID)
	if kaerr != nil {
		fmt.Println(kaerr.Error())
	}
	// 及时从KeepAlive返回的通道中取走值,防止队列充满
	go func() {
		for {
			ka := <-ch
			if ka != nil {
				fmt.Println("ttl:", ka.TTL)
			} else {
				fmt.Println("Unexpected NULL")
			}
		}
	}()

	return nil
}

// GetService 服务发现
func (c *ClientConn) GetService(key string) ([]string, error) {
	resp, err := c.client.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}
	addrs := c.extractAddrs(resp)
	return addrs, nil
}

// extractAddrs 提取地址
func (c *ClientConn) extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}
