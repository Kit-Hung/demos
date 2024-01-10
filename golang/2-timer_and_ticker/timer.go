/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/9 18:18
 * @Description： 定时器示例
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 到点就往通道 C 发送时间值
	timer := time.NewTimer(time.Second)

	// 其他工作 channel
	ch := make(chan int)
	defer close(ch)
	go worker(ch)

	// 为协程设定超时时间
	select {
	case val := <-ch:
		fmt.Printf("revice from ch: %v\n", val)

	case val := <-timer.C:
		fmt.Printf("revice from timer: %v\n", val)
	}

}

func worker(ch chan int) {
	// 工作协程
	time.Sleep(2 * time.Second)
	ch <- 1
}
