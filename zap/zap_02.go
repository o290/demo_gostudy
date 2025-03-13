package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	//logger.Info 是记录一个 INFO 级别日志的方法
	logger.Info("failed to fetch url",
		//使用了 zap 提供的结构化字段
		//这种字段是强类型字段
		zap.String("url", "www.baidu.com"),
		zap.Int("attempt", 1),
		zap.Duration("backoff", time.Second))
	zap.NewDevelopmentEncoderConfig()
}
