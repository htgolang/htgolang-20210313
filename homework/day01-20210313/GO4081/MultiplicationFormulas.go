/**
* @Author: fengzhilaoling
* @Date: 2021/3/14 7:25
**/
package main

import "fmt"

// 打印乘法口诀

func main() {
	fmt.Println("正三角方式打印。。。")
	positiveTriangle()
	fmt.Println("倒三角方式打印。。。")
	invertedTriangle()
}

// 正三角方式打印
func positiveTriangle() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println("")
	}
}

// 倒三角方式打印
func invertedTriangle() {
	for i := 9; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println("")
	}
}
