package main

import "fmt"

func test(flag bool) {
	fmt.Println("call")
	// 看代码的执行逻辑
	if flag {
		return
	}
	fmt.Println("xxx")
}

func main() {
	test(false)
}
