/*
	打印乘法口诀（完成）
    1 * 1 =>1
    1 * 2 => 2 	2 * 2=>4
    1 * 9 => 9 .....    	9*9 => 81
*/

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d\t", i, j, j*i)
		}
		fmt.Println()
	}
}
