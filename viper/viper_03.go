package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	//viper.WriteConfigAs文件存在或不存在都不会报错
	if err := viper.WriteConfigAs("./config/config.yml"); err != nil {
		log.Printf("error:%s", err)
	}
	log.Println("success")
	//viper.SafeWriteConfigAs
	//文件存在不能写入
	if err := viper.SafeWriteConfigAs("./config/config.yml"); err != nil {
		log.Printf("error:%s", err)
	} else {
		log.Println("config.yml success")
	}

	//文件不存在可以写入
	viper.Set("settings.application.host", "0.0.0.0")
	viper.Set("settings.application.name", "write2")
	if err := viper.SafeWriteConfigAs("./config/write2.yml"); err != nil {
		log.Fatalf("error:%s", err)
	} else {
		log.Println("write2.yml success")
	}

}
