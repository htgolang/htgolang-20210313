/*
编程需求：
	打印杨辉三角

*
* *
* * *
* * * *
* * * * *

i:行数 i
j:列数  i

* * * * *
* * * *
* * *
* *
*
*/
/***********************************split line************************/
/*
编程思维：
how print this?

使用for 嵌套循环：
	加判断
	这个和99乘法表应该是类似的；
	braek/continue；
	我先打印一个正方形，可好？
*/
/***********************************split line************************/
/*
开始编程：
*/
package main

import "fmt"

func main() {
	//正序：
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
			if i == j {
				fmt.Println(" ")
			}
		}
	}

	fmt.Println("-----------------------------")

	// //倒序：
	for x := 1; x <= 5; x++ {
		for y := 6 - x; y >= 1; y-- {
			fmt.Printf("* ")
			if y == 1 {
				fmt.Println(" ")
			}
		}
	}
}
