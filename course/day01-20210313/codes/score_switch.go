package main

import "fmt"

func main() {
	/*
		switch {
		case expr:
		case expr1:
		...
		case exprn:

		default:

		}

		expri == true ==> i语句块
	*/

	score := 0

	fmt.Print("请输入成绩：")
	fmt.Scan(&score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("c")
	case score >= 60:
		fmt.Println("d")
	default:
		fmt.Println("e")
	}
}
