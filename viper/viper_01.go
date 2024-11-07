package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		//log.Fatalf() 时，它不仅会输出错误信息，而且会 调用 os.Exit(1)，
		//使程序立即终止并退出，后续的代码（如 fmt.Println()）将不会被执行。
		log.Fatalf("config read error:%s", err)
	}
	log.Println("config read success")

	//设置默认
	viper.SetDefault("settings.application.domain", "localhost:8080")
	domain := viper.GetString("settings.application.domain")
	fmt.Println("Domain:", domain)
	//获取字符串viper.GetString
	host := viper.GetString("settings.application.host")
	fmt.Println("Host:", host)
	//获取布尔值
	ishttps := viper.GetBool("settings.application.ishttps")
	fmt.Println("Is https:", ishttps)
	//获取int
	writertimeout := viper.GetInt("settings.application.writertimeout")
	fmt.Println("writertimeout:", writertimeout)
}
