package main

import (
	"fmt"
)

const version string = "2.2.2"

func main() {
	// 常量必须显示设置默认值
	// 且局部常量必须使用
	const funcVersion string = ""

	const pi = 3.1415926

	const (
		a1     string = "222"
		a2     int    = 32
		a3, a4 string = "1", "2"
	)
	const a5, a6 = "111", 1

	const (
		e1 = "aaa"
		e2 // 只定义标识符 使用上一行的常量标识符的值来进行赋值
		e3
		e4
		e5 = 1
		e6
		e7
	)

	fmt.Println(version, funcVersion, pi, a1, a2, a3, a4, a5, a6)
	fmt.Println(e1, e2, e3, e4, e5, e6, e7)
}
