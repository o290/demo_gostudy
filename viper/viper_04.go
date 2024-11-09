package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func main() {
	//监测和重载配置
	viper.SetConfigName("change")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config error:%s", err)
	}
	log.Println("read config success")

	name := viper.GetString("settings.application.name")
	fmt.Println("name:", name)
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已更改:", e.Name)
	})
	viper.WatchConfig()

	//select让程序一直运行，能够检测到配置文件的变化
	select {}
}
