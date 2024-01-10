/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/10 10:58
 * @Description： 停止子协程示例
 */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. 通过 channel 通知停止
	// 2. 通过 context 超时

	// 通过 channel 通知
	done := make(chan bool)
	go func() {
		select {
		case <-done:
			fmt.Println("child process exits by channel notify...")
		}
	}()

	time.Sleep(time.Second)
	close(done)

	// 通过 context 超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("child process exits by context timeout...")
		}
	}(ctx)
	<-ctx.Done()
	time.Sleep(time.Second)
	fmt.Println("main exits...")
}
