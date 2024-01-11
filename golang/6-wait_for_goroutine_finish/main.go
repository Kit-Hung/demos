/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/10 15:45
 * @Description： 等待一批协程执行完示例
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

const count = 10

func main() {
	// 1. 通过 sleep 等待，缺点： 无法正确评估业务所需时间
	// 2. 通过 channel 判断，缺点： channel 有额外的开销，而且读写数量要协调一致
	// 3. 通过 WaitGroup，推荐做法

	waitBySleep()
	waitByChannel()
	waitByWaitGroup()
}

func waitBySleep() {
	for i := 0; i < count; i++ {
		go fmt.Println("wait by sleep: ", i)
	}

	// 此处无法正确评估业务需要执行多久才能结束
	time.Sleep(5 * time.Second)
}

func waitByChannel() {
	ch := make(chan int, count)
	defer close(ch)

	for i := 0; i < count; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	for i := 0; i < count; i++ {
		fmt.Println("wait by channel: ", <-ch)
	}
}

func waitByWaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Println("wait by WaitGroup: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
