// 3. 猜数字游戏
//     启动程序时生成一个随机数(target)让用户猜测guess(让用户输入数据)
//     猜测太大了 提示 太大了
//     猜测太小了 提示 太小了
//     想等  提示 猜对了 => 程序退出

//     最多猜测5测

/*
编程过程：
随机数如何写？
与用户交互的程序如何写？
init = 原始数据
user = 用户输入的数
最多猜测五次：for循环
选择语句：switch
猜测5次：break


*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	init := rand.Int() % 101
	fmt.Println(init)

	fmt.Println("这是一个猜字游戏，你只有5次机会哦！")
BI:
	for i := 1; i <= 5; i++ {
		user := 0
		fmt.Println("请输入一个数字：")
		fmt.Scan(&user)
		switch {
		case user > init:
			fmt.Println("太大了")
			fmt.Println("---------------------------------")
		case user < init:
			fmt.Println("太小了")
			fmt.Println("---------------------------------")
		default:
			fmt.Println("猜对了")
			fmt.Println("----------------祝贺-----------------")
			break BI
		}
		if i > 5 {
			break BI
		}
	}
}
