package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("write")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config error:%s", err)
	}
	log.Println("read config success")
	// 1.viper.WriteConfig()
	//存在时覆盖，不存在的时候写入
	mode := viper.GetString("settings.application.mode")
	fmt.Println("before mode:", mode)
	viper.Set("settings.application.mode", "release")
	viper.Set("settings.application.port", "8080")
	viper.Set("settings.application.readtimeout", 1)
	viper.WriteConfig()
	mode = viper.GetString("settings.application.mode")
	fmt.Println("after mode:", mode)
	port := viper.GetString("settings.application.port")
	fmt.Println("after port:", port)
	readtimeout := viper.GetString("settings.application.readtimeout")
	fmt.Println("after readtimeout:", readtimeout)
	//2.SafeWriteConfig()
	//存在时不会覆盖，不存在的时候会写入
	viper.Set("settings.application.host", "127.0.0.1")
	viper.Set("settings.application.year", 2024)
	viper.SafeWriteConfig()
	host := viper.GetString("settings.application.host")
	fmt.Println("after host:", host)
	year := viper.GetString("settings.application.year")
	fmt.Println("after year:", year)
	name := viper.GetString("settings.application.name")
	fmt.Println("name:", name)

}
