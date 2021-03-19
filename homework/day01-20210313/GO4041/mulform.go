package main

import "fmt"

func main() {

	// 行 i 列 j 循环
	// 打印乘法口诀的时候，列数《= 行数，打印输出
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", i, j, i*j)
		}

		// 每一行一个换行
		fmt.Println()
	}
}
