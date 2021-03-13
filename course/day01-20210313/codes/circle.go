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
	var confirm string

	// const 定义常量
	const pi float64 = 3.1415926

	for {
		// 从控制台上读取数据
		fmt.Print("请输入半径:") //打印内容后不会加换行
		fmt.Scan(&r)
		// pi = 3.14
		// 计算圆形面积
		fmt.Println(pi * r * r) // 打印内容后会加换行

		// 从控制台接收是否继续的数据
		fmt.Print("是否继续(y/n)?")
		fmt.Scan(&confirm)
		if confirm != "y" {
			break
		}
	}
}
