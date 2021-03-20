package main

import (
	"fmt"
	"math/rand"
	"time"
)

var logo string = `
================猜数游戏=================
			1.最多有五次机会             
			2.所猜数字在(0-100)之间        
=======================================
`

func main() {
	var i, num int
	fmt.Print(logo)
	rand.Seed(time.Now().UnixNano())
	ran := rand.Intn(101)
	for i = 1; i <=5; i++{
		fmt.Printf("请输入您第%d次猜到的数字：", i)
		fmt.Scan(&num)
		if num == ran {
			fmt.Println("恭喜您，猜中了")
			break
		} else if num < ran {
			fmt.Println("您猜的数字小了！！！")
		} else if num > ran {
			fmt.Println("您猜的数字大了！！！")
		} else {
			fmt.Println("输入错误")
		}
	}
}