package init

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func InitEtcd(addr string) *clientv3.Client {
	//配置etcd客户端
	config := clientv3.Config{
		//etcd服务地址列表
		Endpoints: []string{addr},
		//连接超时时间
		DialTimeout: 5 * time.Second,
	}
	// 创建 etcd 客户端实例
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatalf("Failed to create etcd client: %v", err)
	}

	//close也可以不写
	//defer client.Close()
	return client
}
