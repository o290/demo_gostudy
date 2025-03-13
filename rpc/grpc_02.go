package main

import (
	"fmt"
	"net/rpc"
)

// 请求结构体
type Req1 struct {
	Num1 int
	Num2 int
}

// 响应结构体
type Res1 struct {
	Num int
}

func main() {
	//构造请求
	req := Req1{1, 2}
	//建立与服务器的连接
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	//调用远程方法，方法为Add
	var res Res1
	//用client.Call调用
	client.Call("Server.Add", req, &res)
	fmt.Println(res)
}
