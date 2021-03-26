package main

import "fmt"

func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	for i := 0; i < len(slice); i++ {
		// fmt.Println(slice[i])
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[j], slice[i] = slice[i], slice[j]
				// fmt.Println(slice)
				break
			} else if slice[i] <= slice[j] {
				// slice[j], slice[i] = slice[i], slice[j]
				continue
			}
		}
	}
	fmt.Println(slice)
}
