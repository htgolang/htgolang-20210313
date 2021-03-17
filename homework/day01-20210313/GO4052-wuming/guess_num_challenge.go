package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//生成一个50以内的随机数
	var isChoice bool = true

	for isChoice {
		rand.Seed(time.Now().UnixNano())
		var tmp int
		var target int
		tmp = rand.Int()
		target = tmp % 51
		println(target)
		//用户输入一个数字
		var num int

		//判断大小
		for i := 4; i >= 0; i-- {
			println("请输入数字：")
			fmt.Scan(&num)

			if i != 0 {
				switch {
				case num < target:
					fmt.Printf("太小了，你还有%d次机会\n:", i)
				case num > target:
					fmt.Printf("太大了，你还有%d次机会\n:", i)
				case num == target:
					fmt.Printf("猜对了\n")
					break
				default:
					break
				}
			} else {
				fmt.Printf("机会耗尽，正确答案是%d\n", target)
			}

		}
		println("是否继续？(y/n):")
		var choice string
		fmt.Scan(&choice)

		if choice == "y|Y" {
			break
		} else {
			fmt.Println("再见")
			isChoice = false
		}
	}
}