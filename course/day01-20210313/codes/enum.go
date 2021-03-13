package main

import "fmt"

func main() {
	const (
		e1 = iota //0 iota 在小括号内初始化为0，每调用+1
		e2        //1
		e3        //2
	)

	// enum 枚举 星期几 1，2，3，4，5，6，7
	const (
		a1 = iota // 0
		a2        // 1
		a3        // 2
	)
	fmt.Println(e1, e2, e3)
	fmt.Println(a1, a2, a3)
}
