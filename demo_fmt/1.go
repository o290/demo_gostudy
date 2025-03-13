package main

import "fmt"

func main() {
	//输入
	var a1 int
	var a2 int
	var a3 int
	//从标准输入读取数据，当读取到换行符时，函数结束读取。
	fmt.Scan(&a1)
	fmt.Scanln(&a2)
	fmt.Scanf("%d", a3)
	//根据格式化字符串对参数进行格式化，并返回格式化后的字符串，
	//而不直接输出到标准输出
	m := fmt.Sprintf("a1的值为%d", a1)
	n := fmt.Sprint("hello")

	//输出
	fmt.Println(m)
	fmt.Print(n)
	fmt.Printf("a3是%d", a3)

	//根据格式化字符串创建一个新的错误对象并返回。
	//常用于在程序中创建带有自定义错误信息的错误。
	err := fmt.Errorf("错误")
	fmt.Println(err)

}
