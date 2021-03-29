package main

import "fmt"

func main() {
	//第一位数字从1到9
	for i := 1; i<= 9; i++ {
		//第二位数从1到9，每次小于等于第一位数
		for j := 1; j <= i; j++{
			//不能用pringln，因为每打印完会换行
			//格式化输出，结果左对齐占两位
			fmt.Printf("%d * %d = %-2d\t", i, j, i*j)
			//fmt.Println(i, "x", j, "=", i*j)
		}
		//当i=j时候换行
		fmt.Println()
	}
	//fmt.Println()
}