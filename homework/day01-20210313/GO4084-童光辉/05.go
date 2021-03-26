package main

import "fmt"

func main() {
	var num, sum int // 同时定义多个变量，num 键盘接收输入的层数，sum 计算打印的数字
	fmt.Print("请输入杨辉三角的层数：")
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		sum = 1 // 打印每行第一个数字

		for k := num - 1; k > i; k-- {
			fmt.Print(" ") // 打印左边的空格数

		}
		for j := 0; j <= i; j++ {
			fmt.Print(" ", sum)           // 输出 具体的数字
			sum = sum * (i - j) / (j + 1) // 计算每行的数字公示
		}

		fmt.Println()
	}
}
