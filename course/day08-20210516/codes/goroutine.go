package main

import (
	"fmt"
	"time"
)

// 打印A->Z
func printChars(prefix string) {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	// 至少有3个例程
	// c1 c2 交叉执行 同时在执行(并发 只有一个CPU，多个CPU)
	go printChars("c1: ")
	go printChars("c2: ")

	time.Sleep(3 * time.Millisecond) // 加sleep之前没有打印

	// main函数由GO调度器创建例程进行执行 => 主例程
	// 自己用go+函数调用 创建的例程 => 工作例程
	// 当主例程执行完成 自动结束所有工作例程

	// 等待工作例程例程结束后再让主例程退出(主例程要等待(等待某几个，或等待所有的)工作例程结束)
}
