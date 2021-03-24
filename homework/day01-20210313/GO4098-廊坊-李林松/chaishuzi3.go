package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessNum3() {
	var target int
	var input string
	var index = 1
	for {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(11)
		for {
			fmt.Print("请输入十以内的数字:")
			fmt.Scan(&target)
			if target < randNum {
				fmt.Println("输入的数字太小请重新输入")
			} else if target > randNum {
				fmt.Println("输入的数字太大请重新输入")
			} else {
				fmt.Printf("猜对了, 数字是: %d\n", randNum)
				fmt.Print("是否继续游戏(y/n):")
				fmt.Scan(&input)
				if input == "y" {
					break
				} else {
					return
				}
			}
			if index == 5 {
				fmt.Println("最多只能猜5次")
				return
			}
			index++
		}
	}
}
func main() {
	guessNum3()
}

