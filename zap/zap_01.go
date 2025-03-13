package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"time"
)

func main() {
	//zap.NewProduction()创建生产环境配置的zap日志记录器，会产生json格式的日志
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("new production zap fail ")
	}
	//延迟同步日志
	//logger.Sync()用于将日志缓冲区中的日志输出
	//defer在程序结束时将缓存同步到文件中
	defer logger.Sync()
	//创建了一个 sugar 记录器
	//sugar 可以提供带有动态字段的结构化日志，支持Printf格式化日志
	sugar := logger.Sugar()
	url := "https://www.baidu.com/"
	fmt.Println("------------------")
	//infow是一种日志记录方法，用于记录结构化日志
	sugar.Infow("fail to fetch url",
		"url", url,
		"attempt", 3,
		"backofff", time.Second,
	)
	fmt.Println("-----------------")
	//infof是另一种日志记录方法，允许使用格式化字符串来记录日志
	sugar.Infof("fail to fetch url:%s", url)
}
