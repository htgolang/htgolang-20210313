package main

import "fmt"

func main() {

	// 可以通过任意表达式计算结果赋值
	var a1 = 10
	var a2 = a1 + 10
	var a3 = a1 * a2

	// 只能通过字面量
	// 字面量 计算结果
	// 常量的计算结果
	// 或者常量经过某些函数计算的结果
	const e1 = 10
	const e2 = e1 + 10
	const e3 = e2 * e1

	var a4 = a3 + e3

	fmt.Println(a3, e3, a4)
	a4 = 1000
	fmt.Println(a4)
}
