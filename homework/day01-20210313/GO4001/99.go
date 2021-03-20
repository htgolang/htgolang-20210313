package main

import "fmt"

//99乘法表 正三角形
func main() {

	var i, j int
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)

		}
		fmt.Println()
	}
}
