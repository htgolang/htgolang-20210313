package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 申明map
	rep := make(map[string]int)

	// 打开读取文件
	content, err := ioutil.ReadFile("i-have-a-dream.txt")
	if err != nil {
		panic(err)
	}

	// 遍历统计
	for _, i := range string(content) {
		if (i >= 65 && i <= 90) || (i >= 97 && i <= 122) {
			v, ok := rep[string(i)]
			if ok {
				rep[string(i)] = v + 1
			} else {
				rep[string(i)] = 1
			}
		}
	}

	// 输出结果
	for k, v := range rep {
		fmt.Println(k, v)
	}

}
