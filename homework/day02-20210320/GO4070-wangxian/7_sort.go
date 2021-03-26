package main

import "fmt"

func main() {
	var s1 = []int{9, 5, 19, 2, 8, 13, 15, 21, 10, 32}
	// var s1 = []int{1,3,7,5,9}
	for i := 0; i < len(s1)-1; i++ {
		var flag = false
		// fmt.Println("-----")
		for j := 0; j < len(s1)-1-i; j++ {
			if s1[j] > s1[j+1] {
				s1[j], s1[j+1] = s1[j+1], s1[j]
				flag = true
			}
		}
		if !flag{
			break
		}
	}
	fmt.Println(s1)
}
