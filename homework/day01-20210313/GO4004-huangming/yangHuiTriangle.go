package main

import "fmt"

func main() {
	element := map[int][]int{}

	for n := 0; n < 11; n++ {
		if n == 1 {
			fmt.Printf("%d\n", n)
			element[n] = append(element[n], n)
		} else if n == 2 {
			for i := 1; i <= n; i++ {
				fmt.Printf("%-4d", n-1)
				element[n] = append(element[n], n-1)
			}
			fmt.Println()
		} else {
			for i := 1; i <= n; i++ {
				if i == 1 || i == n {
					element[n] = append(element[n], 1)
					fmt.Printf("%-4d", 1)
				} else {
					element[n] = append(element[n], element[n-1][i-2]+element[n-1][i-1])
					fmt.Printf("%-4d", element[n-1][i-2]+element[n-1][i-1])
				}
			}
			fmt.Println()
		}
	}
}
