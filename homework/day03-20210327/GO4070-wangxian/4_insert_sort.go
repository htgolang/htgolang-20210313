package main

import "fmt"

var s1 = []int{5, 7, 9, 1, 3, 11, 2}

func main() {
	for i := 1; i < len(s1); i++ {
		for j := i; j > 0; j-- {
			if s1[j] < s1[j-1] {
				s1[j], s1[j-1] = s1[j-1], s1[j]
			} else {
				break
			}
		}
	}

	fmt.Println(s1)
}
