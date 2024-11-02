package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//写法1
	//basename := "casbin_study"
	a, err := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	//写法2
	//a, err := xormadapter.NewAdapter("mysql", "root:qwer0209@tcp(127.0.0.1:3306)/casbin", true)
	if err != nil {
		log.Fatalf("error:adapter:%s", err)
	}
	m, err := model.NewModelFromString(`
		[request_definition]
		r = sub, obj, act
	
		[policy_definition]
		p = sub, obj, act
	
		[policy_effect]
		e = some(where (p.eft == allow))
	
		[matchers]
		m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("error:model:%s", err)
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error:enforce:%s", err)
	}
	//if err := e.LoadPolicy(); err != nil {
	//	log.Fatalf("failed to load policy: %v", err)
	//}

	// 添加权限
	_, err = e.AddPolicy("alice", "data1", "read")
	if err != nil {
		log.Fatalf("failed to add policy: %v", err)
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
