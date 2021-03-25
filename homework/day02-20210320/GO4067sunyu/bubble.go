package main

import "fmt"

func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	for i := 0; i < len(slice); i++ {
		for k := 0; k < len(slice)-1; k++ {
			if slice[k] > slice[k+1] {
				slice[k], slice[k+1] = slice[k+1], slice[k]
			}
		}
	}
	fmt.Println(slice)
}
