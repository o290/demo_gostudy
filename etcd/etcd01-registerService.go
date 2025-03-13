package main

import (
	"context"
	"demo_gostudy/etcd/init"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

// RegisterService 注册服务到 etcd
func RegisterService(serviceName, serviceAddr string) {
	// 创建一个租约，设置租约时间为 10 秒
	client := init.InitEtcd("127.0.0.1:2379")
	resp, err := client.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatal(err)
	}
	// 服务的键名，格式为 /services/{serviceName}/{serviceAddr}
	key := "/services/" + serviceName + "/" + serviceAddr
	// 将服务地址存储到 etcd 中，并关联租约
	_, err = client.Put(context.TODO(), key, serviceAddr, clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// 保持租约，防止租约过期导致服务信息被删除
	_, err = client.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Service %s registered at %s", serviceName, serviceAddr)
}

func main() {

	// 注册服务
	RegisterService("my-service", "127.0.0.1:8080")

	// 保持程序运行
	select {}
}
