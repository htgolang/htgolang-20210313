package main

import "fmt"

func main() {

	/*
		EXIT:
			for {
				select {
				case expr:
					break EXIT
				}
			}
	*/

	fmt.Println("break:")
BEXIT:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				break BEXIT
			}
			fmt.Println(i, j)
		}
	}
	// 0, 0
	// *1, 0
	// *2, 0

	fmt.Println("continue:")
CEXIT:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue CEXIT
			}
			fmt.Println(i, j)
		}
	}
	//0,0
	//*0,2
	//1,0
	//*1,2
	//2,0
	//*2,2
}
