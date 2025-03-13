package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

// 服务端RPC

// 定义服务端结构体
type Server struct {
}

// 定义请求结构体
type Req struct {
	Num1 int
	Num2 int
}

// 定义响应结构体
type Res struct {
	Num int
}

// 定义服务端的方法
func (s Server) Add(req Req, res *Res) error {
	res.Num = req.Num1 + req.Num2
	fmt.Println("response is ", req)
	return nil
}

func main() {
	// 注册rpc服务
	rpc.Register(new(Server))
	//将RPC服务绑定到Http服务中，即服务端可以通过HTTP协议实现远程调用
	rpc.HandleHTTP()
	//监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	//启动HTTP服务
	http.Serve(listen, nil)
}
