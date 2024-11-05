package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

// api的使用
func main() {
	e, err := casbin.NewEnforcer("conf/01.conf", "conf/01.csv")
	if err != nil {
		log.Fatalf("error:enforcer:%s", err)
	}

	//获取policy
	p, _ := e.GetPolicy()
	fmt.Println("policy:", p)
	//添加policy
	added, _ := e.AddPolicy("bob", "data3", "read")
	p, _ = e.GetPolicy()
	fmt.Println(added)
	fmt.Println("policy:", p)
	//修改policy
	updated, err := e.UpdatePolicy([]string{"bob", "data3", "read"}, []string{"bob", "data3", "write"})
	p, _ = e.GetPolicy()
	fmt.Println(updated)
	fmt.Println("policy:", p)
	//删除policy
	removed, _ := e.RemovePolicy("bob", "data3", "read")
	fmt.Println(removed)
	removed, _ = e.RemovePolicy("bob", "data3", "write")
	fmt.Println(removed)
	p, _ = e.GetPolicy()
	fmt.Println("policy:", p)

}
