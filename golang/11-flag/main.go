/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/30 21:58
 * @Description： flag 使用示例
 */
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "the name you want to say hi")
	flag.Parse()

	fmt.Println("os args are: ", os.Args)
	fmt.Printf("hello %s\n", *name)
}
