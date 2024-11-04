package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	//使用模型文件和默认的文件适配器

	//第一种定义方式
	//这个路径来源于casbin github：https://github.com/casbin/casbin/tree/master/examples
	//https://casbin.org/zh/docs/adapters/
	////NewAdapter用于创建适配器
	//a := fileadapter.NewAdapter("conf/01.csv")
	////NewEnforcer用于创建casbin
	//e, err := casbin.NewEnforcer("conf/01.conf", a)
	//if err != nil {
	//	log.Fatalf("error:enforcer:%s", err)
	//}

	//第二种定义方式
	e, err := casbin.NewEnforcer("conf/01.conf", "conf/01.csv")
	if err != nil {
		log.Fatalf("error:enforcer:%s", err)
	}
	sub := "alice"
	obj := "data1"
	act1 := "read"
	act2 := "write"
	//Enforce是一个用于执行权限检查的重要方法，根据定义的访问控制模型和策略，判断某个主体是否对某个资源执行特定操作的权限。
	res1, err := e.Enforce(sub, obj, act1)
	if err != nil {
		fmt.Println("alice can not read data1")
		log.Fatalf("error enforcing policy: %s", err)
	}
	log.Printf("Alice can read data1: %v", res1)

	res2, err := e.Enforce(sub, obj, act2)
	if err != nil {
		fmt.Println("alice can not write data1")
		log.Fatalf("error enforcing policy: %s", err)
	}
	log.Printf("Alice can write data1: %v", res2)
}
