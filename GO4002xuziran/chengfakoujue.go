//打印乘法口诀 1 1 =>1 2 * 1 => 2 2 * 2=>4 9 * 1 => 9 ..... 9*9 => 81

package main

import "fmt"

func main() {
	for n := 1; n < 10; n++ {
		for m := 1; m <= n; m++ {
			fmt.Printf("%d * %d = %d,", m, n, (m * n))
		}
		fmt.Println()

	}
}
