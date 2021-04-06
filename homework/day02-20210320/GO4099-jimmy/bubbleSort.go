package main

import "fmt"

func main() {
	values := []int{9, 8, 19, 10, 2, 8}

	for i := 0; i < len(values)-1; i++ {
		for j := 1 + i; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println(values)
}
