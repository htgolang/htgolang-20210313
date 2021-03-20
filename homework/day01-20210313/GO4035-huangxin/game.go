package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	descir: 猜数游戏
	vesion: 1.0
	authro: huangxin
	date:   2021-03-14
*/
func main() {

	rand.Seed(time.Now().UnixNano()) // 设置随机数种子

	targetNum := rand.Intn(101) // 生成1~100的随机书

	guessNum := 0 // 定义变量接收用户输入
	var i = 5
	for ; i > 0; i-- {
		fmt.Print("请输入一个数:")
		fmt.Scan(&guessNum)

		if guessNum == targetNum {
			fmt.Println("猜对了，Game Over！")
			break
		} else if guessNum > targetNum {
			fmt.Println("大了点，请继续猜！")
		} else {
			fmt.Println("小了点，请继续猜！")
		}
	}
	// 玩游戏失败提示
	if i == 0 {
		fmt.Println("游戏次数用完，Game Over！")
	}

}
