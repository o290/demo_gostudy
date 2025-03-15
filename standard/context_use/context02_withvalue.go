package main

import (
	"context"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

//使用WithValue实现日志跟踪

// 定义一个键类型，用于在 context 中存储跟踪 ID
// 以struct作为键更加节省内存
type traceIDKey struct{}

// 生成一个唯一的跟踪 ID
func generateTraceID() string {
	//生成uuid
	id, err := uuid.NewRandom()
	if err != nil {
		return "unknown"
	}
	return id.String()
}

// 添加跟踪 ID 到 context
func addTraceIDToContext(ctx context.Context) context.Context {
	traceID := generateTraceID()
	//func WithValue(parent Context, key, val any) Context {
	//WithValue相当于创建了一个携带键值对信息的新的context，新的context会继承parent
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// 从 context 中获取跟踪 ID
func getTraceIDFromContext(ctx context.Context) string {
	//ctx.Value(traceIDKey{})从上下文中获取指定键的值
	//.(string)进行类型断言
	traceID, ok := ctx.Value(traceIDKey{}).(string)
	if !ok {
		return "unknown"
	}
	return traceID
}

// 模拟一个处理函数
func processRequest(ctx context.Context) {
	traceID := getTraceIDFromContext(ctx)
	log.Printf("[%s] Processing request...", traceID)
	// 模拟处理时间
	time.Sleep(3 * time.Second)
	log.Printf("[%s] Request processed.", traceID)
}

func main() {
	var wg sync.WaitGroup

	// 模拟多个请求
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := addTraceIDToContext(context.Background())
			processRequest(ctx)
		}()
	}

	wg.Wait()
}
