package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Server struct {
	Port string
	Host string
	Name string
}

func main() {
	viper.SetConfigName("config2")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("read config fail")
	}
	f := func(v *viper.Viper) *Server {
		return &Server{
			Port: v.GetString("port"),
			Host: v.GetString("host"),
			Name: v.GetString("name"),
		}
	}
	vServer := viper.Sub("application.server")
	if vServer == nil {
		log.Fatalln("not find server config")
	}
	serverConfig := f(vServer)
	fmt.Println(serverConfig)
}
