package main

import "context"

func main() {
	//两种创建context的方式
	//这两种方式都是在创建根context，不具备任何功能，功能实现需要通过with系列函数
	context.Background()
	context.TODO()

}
