package main

import (
	"fmt"
)

func main() {
	var slice []int = []int{9, 8, 19, 10, 2, 8}
	var num int
	var index int
	for k,v := range slice {
		if v > num {
			num = v
			index = k
		}else {
			continue
		}
	}
	fmt.Println("The Max is:", num, "The index is:", index)

	
}