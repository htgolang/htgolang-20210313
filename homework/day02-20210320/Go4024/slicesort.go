package main

import (
	"fmt"
)

func main() {
	var slice []int = []int{9, 8, 19, 10, 2, 8}
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j< len(slice)-1-i;j++ {
			if slice[j] > slice[j+1] {
				slice[j],slice[j+1] = slice[j+1],slice[j]
			}
		}
	}
	fmt.Println("排序后为:",slice)
}