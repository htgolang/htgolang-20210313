package main

import "fmt"

func main() {
	s := []int{9, 8, 19, 10, 2, 8}
	max := s[0]
	second := 0
	for i := 0; i < len(s)-1; i++ {
		if max < s[i+1] {
			max = s[i+1]
			second = s[i]
		} else {
			if max > second {
				if second < s[i+1] {
					second = s[i+1]
				}
			}
		}
	}

	for i, v := range s {
		if v == second {
			fmt.Println(i, second)
		}
	}
}
