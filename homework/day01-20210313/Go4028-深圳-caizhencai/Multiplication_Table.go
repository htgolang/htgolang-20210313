package main

/*
打印乘法口诀（完成） 1 1 =>1 2 * 1 => 2 2 * 2=>4 9 * 1 => 9 ..... 9*9 => 81
*/
import "fmt"

func main() {
	for i := 1; i <= 9; i++ {

		for j := 1; j <= i; j++ {
			//var sum int = 0
			//sum = i * j
			if j == i {
				fmt.Printf("%d + %d = %d \n", i, j, i*j)
			} else {
				fmt.Printf("%d + %d = %d \t", i, j, i*j)
			}

		}
	}
}
