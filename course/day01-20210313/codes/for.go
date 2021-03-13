package main

import "fmt"

func main() {
	/*
		for init; expr; end {

		}

		// init
		// expr => true 循环体
		// end -> expr -> 循环体-> end
	*/

	// 1 + 2 + ... + 100
	var total int = 0
	for i := 1; i <= 100; i++ {
		total += i
	}
	fmt.Println(total)

	total = 0
	var i int = 1
	for i <= 100 {
		total += i
		i++
	}
	fmt.Println(total)

	// while(true) {}
	var txt string
	for {
		fmt.Print("请输入内容：")
		fmt.Scan(&txt)
		fmt.Println("你的输入为:", txt)
		if txt == "exit" {
			break
		}
	}
}
