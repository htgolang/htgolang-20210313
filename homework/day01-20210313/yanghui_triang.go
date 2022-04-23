package main

import "fmt"

func main() {
	const num int = 10
	var arr [num][num]int
	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			if i < 2 {
				arr[i][j] = 1
			} else if j == 0 || j == i {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
			}
			fmt.Printf("%d\t", arr[i][j])
		}
		// fmt.Println(arr[i][:i+1])
		fmt.Printf("\n")
	}
	// fmt.Println(arr)
}
