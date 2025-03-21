package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker is canceled:", ctx.Err())
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 创建一个可取消的上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 启动一个工作 goroutine
	go worker(ctx)

	// 模拟一段时间后取消上下文
	time.Sleep(3 * time.Second)
	cancel()

	// 等待一段时间，确保工作 goroutine 有足够的时间响应取消信号
	time.Sleep(1 * time.Second)
	fmt.Println("Main function is exiting.")
}
