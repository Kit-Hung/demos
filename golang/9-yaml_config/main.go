/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/30 17:51
 * @Description： 读 yaml 配置文件示例
 */
package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	A int      `yaml:"a"`
	B string   `yaml:"b"`
	C bool     `yaml:"c"`
	D []string `yaml:"d"`
	E struct {
		EA string `yaml:"ea"`
		EB string `yaml:"eb"`
	}
}

func main() {
	// 读取指定文件
	filename := "config.yaml"
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// 解析文件内容
	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	// 对 config 进行过编码
	yml, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", config)
	fmt.Println("================================")
	fmt.Printf("%s\n", yml)
}
