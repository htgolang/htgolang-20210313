package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(101)
	fmt.Println(num)

	var guess_num int

	for count := 0; count < 5; count++ {
		fmt.Print("请输入你才的数字(0-100)：")
		fmt.Scanln(&guess_num)

		switch {
		case guess_num > num:
			fmt.Println("猜大了")
		case guess_num < num:
			fmt.Println("猜小了")
		default:
			fmt.Println("正确")
			os.Exit(0)
		}
	}
}
