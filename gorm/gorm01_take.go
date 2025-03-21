package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//安装
// go get -u github.com/go-sql-driver/mysql
// go get gorm.io/gorm
// go get gorm.io/driver/mysql

// 创建数据库
// mysql -u root -p
// create database name

type UserConf struct {
	ID     int
	UserID int
	Name   string
}

func main() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/im_server?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:qwer0209@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	//打开数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("格式错误%s", err)
	}
	db.AutoMigrate(&UserConf{})
	// 定义一个变量来存储查询结果
	var userConf UserConf
	//Take查询返回第一条符合条件的记录
	err = db.Take(&userConf, "user_id=?", 3).Error
	if err != nil {
		fmt.Printf("查询失败：%s", err)
	}
	fmt.Println("查询结果：", userConf)
}
