//9*9乘法口诀
package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ", i, j, i*j)
			if i == j {
				fmt.Println(" ")
			}
		}
	}
}

// fmt.Printf("%v*%v=", i, j)

// 1 * 1 = 1
// 2 * 1 = 2 2 * 2 = 4
// 3 * 1 = 3 3 * 2 = 6 3 * 3 =9

// ……
// 9 * 1 = 9……				9 * 9 = 81

// i * j = ji
//    i <= j

//for 循环
// if else
// i j 变量
