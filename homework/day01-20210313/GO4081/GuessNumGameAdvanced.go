/**
* @Author: fengzhilaoling
* @Date: 2021/3/14 8:18
**/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机数游戏进阶
/*
启动程序时生成一个随机数(target)让用户猜测guess(让用户输入数据)
猜测太大了 提示 太大了
猜测太小了 提示 太小了
想等  提示 猜对了 => 询问是否继续 -> 继续 重新生成随机数 不继续退出

最多猜测5测
*/
func guessNumberGame() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	fmt.Println("------欢迎来到猜数字游戏------")

CONTINUEGAME:
	// 1-100之间的随机数
	random := rand.Intn(100) + 1
	fmt.Println(random)

	for i := 5; i > 0; i-- {
		var num int
		fmt.Printf("你有%d次输入机会\n", i)
		fmt.Print("请输入一个1-100这就的数字：")
		fmt.Scanf("%d\n", &num)

		if num > random {
			fmt.Println("你输入的数字太大了")
		} else if num < random {
			fmt.Println("你输入的数字太小了")
		} else {
			fmt.Println("恭喜你输入的数字正确！！！")
			var str string
			fmt.Print("是否重新开始新一轮游戏(y/n)：")
			fmt.Scanf("%s\n", &str)

			if str == "y" {
				fmt.Println("--------开始新的一轮游戏--------")
				goto CONTINUEGAME
			} else if str == "n" {
				fmt.Println("游戏退出！！！")
				break
			}
			fmt.Println("你输入的指令不正确游戏退出！！！")
			break
		}
		if i == 1 {
			fmt.Println("你的5次机会已经用完，游戏退出！！！")
			break
		}
	}
}

func main() {
	guessNumberGame()
	const (
		a int = iota
	)
}
