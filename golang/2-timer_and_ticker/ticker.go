/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/9 18:52
 * @Description： 节拍器示例
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 以指定的时间间隔重复向通道 C 发送时间值
	ticker := time.NewTicker(time.Second)

	for val := range ticker.C {
		fmt.Printf("receive from ticker: %v\n", val)
	}
}
