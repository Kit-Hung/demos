/*
 * @Author:       Kit-Hung
 * @Date:         2024/5/10 11:13
 * @Description： 带超时控制的同步 map
 */
package main

import (
	"context"
	"sync"
	"time"
)

/*
实现一个 map
1. 面向高并发
2. 只存在插入和查询操作 O(1)
3. 查询时，若 key 存在，直接返回 val ，若 key 不存在，阻塞直到 key val 对被放入后，获取 val 返回，等待指定时长仍未放入，返回超时错误

*/

type MyConcurrentMap struct {
	sync.Mutex
	mp      map[int]int
	keyToCh map[int]*MyChannel
}

type MyChannel struct {
	sync.Once
	ch chan struct{}
}

func NewMyChannel() *MyChannel {
	return &MyChannel{
		ch: make(chan struct{}),
	}
}

func (c *MyChannel) Close() {
	c.Do(func() {
		close(c.ch)
	})
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		mp:      make(map[int]int),
		keyToCh: make(map[int]*MyChannel),
	}
}

func (m *MyConcurrentMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v

	ch, ok := m.keyToCh[k]
	if !ok {
		// 不存在说明没有在等待读的 ch
		return
	}

	// 通过 close 通知所有在等待读的 ch，不能重复关闭
	ch.Close()
}

func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	if val, ok := m.mp[k]; ok {
		m.Unlock()
		return val, nil
	}

	ch, ok := m.keyToCh[k]
	if !ok {
		ch = NewMyChannel()
		m.keyToCh[k] = ch
	}
	m.Unlock()

	tCtx, cancel := context.WithTimeout(context.Background(), maxWaitingDuration)
	defer cancel()

	select {
	case <-tCtx.Done():
		// 超时时返回超时错误
		return -1, tCtx.Err()
	case <-ch.ch:
	}

	m.Lock()
	val := m.mp[k]
	m.Unlock()
	return val, nil
}
