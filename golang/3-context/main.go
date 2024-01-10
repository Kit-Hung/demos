/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/9 18:37
 * @Description： 上下文相关示例
 */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// context.Background: 一般作为顶层 context
	// context.TODO: 不确定使用什么的时候使用
	// context.WithDeadline: 可以设置超时时间
	// context.WithValue: 向 context 添加键值对
	// context.WithCancel: 可取消

	// 创建一个顶层 context
	baseCtx := context.Background()

	// 创建一个设置了键值对的 context
	ctx := context.WithValue(baseCtx, "a", "b")
	// 取出对应的值
	go func(ctx context.Context) {
		fmt.Println(ctx.Value("a"))
	}(ctx)

	// 创建一个有超时控制的 context
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()

	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default...")
			}
		}

	}(timeoutCtx)

	time.Sleep(2 * time.Second)
	select {
	case <-timeoutCtx.Done():
		time.Sleep(time.Second)
		fmt.Println("main process exits...")
	}
}
