package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a string
	fmt.Println("请输入数字：：")
	fmt.Scan(&a)
	if id, err := strconv.Atoi(a); err == nil {

		fmt.Println(id)

	} else {

		fmt.Println("hello")
	}

}
