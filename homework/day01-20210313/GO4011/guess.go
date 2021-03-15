package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	启动程序时生成一个随机数(target)，让用户猜测(guess)
    猜大 -> 提示 太大了
    猜小 -> 提示 太小了
    相同 -> 提示 猜对了
*/

func main() {
	// 使用时间戳
	rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Int())

	// 将数值控制在0~100之间
	var target = (rand.Int() % 100)
	fmt.Println(target)

	// 定义让用户输入的变量
	var guessNum int

	for i := 0; i < 5; i++ {
		fmt.Println("请输入0~100之间的整数：")
		fmt.Scan(&guessNum)
		if guessNum == target {
			fmt.Println("猜对了！！")
			break
		} else if guessNum > target {
			fmt.Println("太大了！")
		} else if guessNum < target {
			fmt.Println("太小了！")
		}
	}
	fmt.Println("Game Over")

}
