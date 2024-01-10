/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/9 18:16
 * @Description： 生产者、消费者、单向通道示例
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建双向数据交互通道
	messages := make(chan int, 10)
	defer close(messages)

	// 创建程序退出信号通道, 通知子协程退出
	done := make(chan bool)

	go prod(messages)
	go consume(messages, done)

	// 5 秒之后让程序退出
	time.Sleep(5 * time.Second)
	close(done)

	// 再等 1 秒
	time.Sleep(1 * time.Second)
	fmt.Println("main process exits...")
}

func prod(messages chan<- int) {
	// 生产者使用只发送通道
	for i := 0; i < 10; i++ {
		messages <- i
	}
}

func consume(messages <-chan int, done <-chan bool) {
	// 消费者使用只接收通道
	// 使用节拍器每 1 秒获取一次数据
	ticker := time.NewTicker(time.Second)

	for _ = range ticker.C {
		select {
		case <-done:
			fmt.Println("process interrupt...")
			return
		default:
			val, notClosed := <-messages
			fmt.Printf("revice from channel, value: %v, notClosed: %v\n", val, notClosed)
		}
	}
}
