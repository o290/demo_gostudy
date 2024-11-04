package main

import (
	"github.com/casbin/casbin/v2"
	"log"
)

// 自定义函数在参数中使用

// 1.定义函数，函数的返回值为布尔类型
func match(key1 string, key2 string) bool {
	if key1 == key2 {
		return true
	}
	return false
}

// 2.使用interface包装
func matchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(match(name1, name2)), nil
}
func main() {
	//1.创建casbin适配器
	e, err := casbin.NewEnforcer("conf/03.conf", "conf/03.csv")
	if err != nil {
		log.Fatalf("error enforcer:%s", err)
	}
	//2.将函数注册到casbin执行器中
	//my_func指定方法名字
	e.AddFunction("my_func", matchFunc)
	//3.检查判断
	sub1 := "alice"
	obj1 := "data1"
	act1 := "read"
	res1, err1 := e.Enforce(sub1, obj1, act1)
	if err1 != nil {
		log.Fatalf("error:%s", err1)
	}
	log.Printf("Alice can read data1: %v", res1)
	sub2 := "bob"
	obj2 := "data2"
	act2 := "read"
	res2, err2 := e.Enforce(sub2, obj2, act2)
	if err2 != nil {
		log.Fatalf("error:%s", err2)
	}
	log.Printf("Bob can read data2: %v", res2)
}
