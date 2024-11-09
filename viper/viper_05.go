package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("name")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config fail:%s", err)
	}
	log.Println("read config success")
	//别名RegisterAlias，第一个参数是原来的名字，第二个参数是别名
	viper.RegisterAlias("name", "project")
	viper.Set("project", "project") //用别名设置配置项的值，会覆盖
	fmt.Println(viper.GetString("name"))

	//提取子结构
	son1 := viper.Sub("settings.application.parent.son1")
	if son1 == nil { // 如果不存在返回nil
		log.Fatalln("son1 is not exits")
	}
	fmt.Println(son1.GetString("name"))
	son2 := viper.Sub("settings.application.parent.son2")
	if son1 == nil { // 如果不存在返回nil
		log.Fatalln("son2 is not exit")
	}
	fmt.Println(son2.GetString("name"))

	if son3 := viper.Sub("settings.aplication.parent.son3"); son3 == nil {
		log.Fatalln("son3 is not exit")
	}
}
