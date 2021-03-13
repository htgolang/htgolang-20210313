package main

import "fmt"

// 定义全局变量
// 显示赋值
var version string = "1.1.1"

func main() {
	// go中要求局部变量定义后必须使用
	// 定义局部变量
	// 为对变量进行显式赋值，go会对变量设置默认值（零值）
	// ""
	var funcVersion string

	// 省略类型，必须显示设置默认值，go根据字面量或赋值值类型推导标识符类型
	var name = "kk"
	// var name string = "kk"

	var (
		a1      int
		a2      string
		a3      int    = 3 + 2
		a4      string = "aaa"
		a5             = a3 + a1
		a6             = "bbbb"
		a9, a10 int
	)
	var a7, a8 string
	var a11, a12 string = "a11", "a12"

	fmt.Println(funcVersion, version, name,
		a1, a2, a3, a4, a5, a6, a7, a8, a9, a10,
		a11, a12)
}
