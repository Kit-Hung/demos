/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/10 15:34
 * @Description： 锁使用示例
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go lock()
	go rLock()
	go wLock()

	time.Sleep(5 * time.Second)
}

func lock() {
	// 普通锁
	lock := sync.Mutex{}

	for i := 0; i < 3; i++ {
		lock.Lock()
		// 注意： 此处要到函数返回才会执行，所以下一个循环只会卡住拿不到锁
		defer lock.Unlock()

		fmt.Println("lock: ", i)
	}
}

func rLock() {
	// 读锁
	lock := sync.RWMutex{}

	for i := 0; i < 3; i++ {
		lock.RLock()
		// 因为读锁是共享的，此处不会因为拿不到锁阻塞
		defer lock.RUnlock()

		fmt.Println("rLock: ", i)
	}
}

func wLock() {
	// 写锁
	lock := sync.RWMutex{}

	for i := 0; i < 3; i++ {
		lock.Lock()
		// 注意： 此处要到函数返回才会执行，所以下一个循环只会卡住拿不到锁
		defer lock.Unlock()

		fmt.Println("wLock: ", i)
	}
}
