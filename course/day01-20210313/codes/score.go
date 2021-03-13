package main

import "fmt"

func main() {
	// 从控制台让用户输入一个成绩
	// >= 90 => A
	// >= 80 => B
	// >= 70 => C
	// >= 60 => D
	// < 60 => E

	var score int
	fmt.Print("请输入成绩")
	fmt.Scan(&score)

	// else if
	/*
		if expr1 {

		} else if expr2 {

		} else if expr3 {

		....
		} else if exprn {

		} else {

		}
		expr1 => true if => expr2
		expr2 => true 2语句块 => expr3
		exprn => true n语句块 => else
		if 必须有
		else if 任意多个
		else 0个或1个

	*/
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
