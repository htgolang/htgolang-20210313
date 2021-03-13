package main

import "fmt"

func main() {
	// var xxx int

	// 简短声明 ，通过值来推导标识符类型
	// 只能用在局部
	// 多种数据类型使用同一个字面量标识 只能推导为默认类型
	// 整数 byte int int8 int16 int32 int64 uint
	name := "kk"
	id := 1

	fmt.Println(name, id)
}
