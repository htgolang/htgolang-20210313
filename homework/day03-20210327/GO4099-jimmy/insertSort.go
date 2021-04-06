package main

import "fmt"

func main() {

	s := []int{23, 0, 12, 56, 34}

	for i := 1; i < len(s); i++ {
		insertVal := s[i]
		insertIndex := i - 1

		for insertIndex >= 0 {
			if s[insertIndex] < insertVal {
				s[insertIndex], s[insertIndex+1] = s[insertIndex+1], s[insertIndex]
				insertIndex--
			} else {
				break
			}
		}
	}
	fmt.Println(s)
}
