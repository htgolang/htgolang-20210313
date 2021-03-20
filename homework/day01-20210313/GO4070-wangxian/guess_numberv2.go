package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func rand_num(max int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	num := rand.Intn(max)
// 	fmt.Println(num)
// 	return num
// }

func main() {
START:
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(101)
	fmt.Println(num)

	var guess_num int

	for count := 0; count < 5; count++ {
		fmt.Print("请输入你猜的数字：")
		fmt.Scanln(&guess_num)

		if guess_num > num {
			fmt.Println("猜大了")
		} else if guess_num < num {
			fmt.Println("猜小了")
		} else {
			var confirm string
			fmt.Print("猜对了，还要继续猜吗(y/n):")
			fmt.Scanln(&confirm)

			if confirm == "y" {
				goto START
			} else {
				fmt.Println("再见")
				break
			}
		}
		fmt.Printf("还剩余%d次机会\n", 4-count)
	}
}
