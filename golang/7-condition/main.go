/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/10 16:06
 * @Description： 通过 sync.Cond 实现队列，做生产者消费者模型
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	// 启动生产者
	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(2 * time.Second)
		}
	}()

	// 消费者
	for {
		result := q.Dequeue()
		fmt.Println("dequeue result: ", result)
		time.Sleep(time.Second)
	}
}

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func (q *Queue) Enqueue(item string) {
	// 生产者，数据进队列
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.queue = append(q.queue, item)
	fmt.Printf("putting item %v to queue, notify all...\n", item)
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() string {
	// 消费者，数据出队列
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if len(q.queue) == 0 {
		fmt.Println("no data available, wait...")
		q.cond.Wait()
	}

	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}
