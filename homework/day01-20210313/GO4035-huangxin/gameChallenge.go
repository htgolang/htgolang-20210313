package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子

	targetNum := rand.Intn(101) // 生成1~100的随机书
	fmt.Println(targetNum)
	guessNum := 0 // 定义变量接收用户输入

	var i = 5 // 记录游戏次数

	var ifContinue string // 游戏是否继续(y:继续)

	for ; i > 0; i-- {
		fmt.Print("请输入一个数:")
		fmt.Scan(&guessNum) // 键盘接收用户输入的数字

		if guessNum == targetNum {
			fmt.Println("太棒了！恭喜你猜对了！")
			fmt.Print("您是否还想再玩一次(y/n):")
			fmt.Scan(&ifContinue) // 键盘输入是否继续玩游戏

			if ifContinue == "y" {
				i = 5                      // 将游戏次数重置
				targetNum = rand.Intn(101) // 将随机数重置，生成新的随机数
				continue
			} else {
				fmt.Println("游戏退出，期待您下次再来！")
				break
			}
		} else if guessNum > targetNum {
			fmt.Println("大了点，请继续猜！")
		} else {
			fmt.Println("小了点，请继续猜！")
		}
	}
	// 玩游戏失败的提示
	if i == 0 {
		fmt.Println("游戏次数用完，Game Over！")
	}
}
