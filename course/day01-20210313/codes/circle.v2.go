package main

import (
	"fmt"
)

func main() {
	// var 定义变量
	// 变量名称 r
	// 变量类型 float64
	// r的值为10
	var r float64 = 100

	// const 定义常量
	const pi float64 = 3.1415926

	r = 5
	// pi = 3.14
	// 计算圆形面积
	fmt.Println(pi * r * r)
}
