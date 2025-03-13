package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// DiscoverServices 从 etcd 中发现服务
func DiscoverServices(client *clientv3.Client, serviceName string) ([]string, error) {
	// 构造服务的键前缀，用于查询该服务下的所有实例
	keyPrefix := "/services/" + serviceName + "/"
	// 从 etcd 中获取所有匹配该前缀的键值对
	resp, err := client.Get(context.TODO(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var serviceAddrs []string
	// 遍历查询结果，将服务地址添加到切片中
	for _, ev := range resp.Kvs {
		serviceAddrs = append(serviceAddrs, string(ev.Value))
	}

	return serviceAddrs, nil
}

func main() {
	// 连接到 etcd 服务
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 发现服务
	addrs, err := DiscoverServices(cli, "my-service")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Discovered services: %v", addrs)
}
