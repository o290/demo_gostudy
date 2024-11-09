package main

import (
	"fmt"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
	"log"
)

type config struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
}

func main() {
	viper.SetConfigName("config2")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config fail:%v", err)
	}

	//反序列化至结构体
	var serverConfig config
	if err := viper.UnmarshalKey("application.server", &serverConfig); err != nil {
		log.Fatalf("Error unmarshaling key, %s", err)
	}
	fmt.Println(serverConfig)

	//反序列化至map
	var dbConfig map[string]interface{}
	var config2 map[string]map[string]interface{}
	if err := viper.UnmarshalKey("application.database", &dbConfig); err != nil {
		log.Fatalf("Error unmarshaling key, %s", err)
	}
	fmt.Println(dbConfig)
	if err := viper.Unmarshal(&config2); err != nil {
		log.Fatalf("Error unmarshaling, %s", err)
	}
	fmt.Println(config2)
	for _, v := range config2 {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	//序列化
	yamlData, err := yaml.Marshal(&serverConfig)
	if err != nil {
		log.Fatalf("Error marshaling to YAML, %s", err)
	}
	fmt.Printf("Serialized YAML:\n%s\n", string(yamlData))
	yamlData2, err := yaml.Marshal(&dbConfig)
	if err != nil {
		log.Fatalf("Error marshaling to YAML, %s", err)
	}
	fmt.Printf("Serialized YAML:\n%s\n", string(yamlData2))
}
