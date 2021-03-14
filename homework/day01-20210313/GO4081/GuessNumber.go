/**
* @Author: fengzhilaoling
* @Date: 2021/3/14 7:44
**/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 猜数字游戏
/*
启动程序时生成一个随机数(target)让用户猜测guess(让用户输入数据)
猜测太大了 提示 太大了
猜测太小了 提示 太小了
想等  提示 猜对了 => 程序退出

最多猜测5测
*/
func guessNumGame() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	// 1-100之间的随机数
	random := rand.Intn(100) + 1
	//fmt.Println(random)
	var num int
	// 计算用户输入的次数，超过5次退出
	var count int = 5
	fmt.Printf("------欢迎来到猜数字游戏------\n")
	for {
		fmt.Printf("你有%d次输入机会。。。\n", count)
		fmt.Print("请输入一个1-100之间的数字：")
		fmt.Scanf("%d\n", &num)
		if random > num{
			fmt.Println("你输入的数字太小了")
		}else if random < num {
			fmt.Println("你输入的数字太大了")
		}else {
			fmt.Println("恭喜你输入的数字正确！")
			fmt.Println("游戏退出。。。")
			break
		}
		count--
		if count == 0{
			fmt.Println("你的5次机会已经用完，游戏失败！！！")
			break
		}
	}

}

func main() {
	guessNumGame()
}
