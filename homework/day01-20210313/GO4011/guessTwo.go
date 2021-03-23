package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 使用时间戳作为随机数
	rand.Seed(time.Now().UnixNano())
	/*
		// 截取两值数0-99
		var target = (rand.Int() % 100)
		fmt.Println(target)
	*/

	confirm := ""

	// 设置让用户最多猜五次，完成后让用户确认是否继续
	a := 0
	for {
		a++

		// 截取两值数0-99
		var target = (rand.Int() % 100)
		fmt.Println(target)

		// 定义用户输入数据的变量
		var guessNum int
		for i := 1; i <= 5; i++ {
			// 用户输入数字
			fmt.Println("请输入1~100之间的整数：")
			fmt.Scan(&guessNum)

			// 用户输入正确 => 猜对； 用户输入大 => 太大了；用户输入小 => 太小了
			// 5次后让用户确认是否继续。
			if guessNum == target {
				fmt.Println("正确")
				break
			} else if guessNum > target {
				fmt.Println("太大了")
			} else if guessNum < target {
				fmt.Println("太小了")
			}
			//break

		}
		fmt.Println("Game Over")

		// 询问用户是否继续
		fmt.Println("是否继续：(y/n)")
		fmt.Scan(&confirm)
		if confirm == "y" {
			fmt.Println("Welcome !")
		} else {
			break
		}
	}
}
