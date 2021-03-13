package main

import "fmt"

func main() {
	/*
		switch value {
		case tvalue1:
		case tvalue2, tvalue3:
		...
		case tvaluen:

		default:

		}

		value == tvaluei ==> i语句块
	*/

	confirm := ""
	fmt.Println("老婆想法：")
	fmt.Println("十个包子")

	fmt.Print("有卖西瓜的吗?")
	fmt.Scan(&confirm)

	// confirm == "y"

	switch confirm {
	case "y":
		fmt.Println("一个西瓜")
	}

	fmt.Println("老公想法：")
	fmt.Print("有卖西瓜的吗?")
	fmt.Scan(&confirm)

	// confirm == "y"

	switch confirm {
	case "y":
		fmt.Println("一个包子")
	default:
		fmt.Println("10个包子")
	}
}
