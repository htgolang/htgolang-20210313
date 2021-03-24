/*
	猜数字游戏
    启动程序时生成一个随机数(target)让用户猜测guess(让用户输入数据)
    猜测太大了 提示 太大了
    猜测太小了 提示 太小了
    相等  提示 猜对了 => 程序退出
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const input_count = 5

func main() {
	var input_num int

	rand.Seed(time.Now().Unix()) //使用当前时间设置随机数种子

	result := rand.Int() % 101 //生成[0,101)的随机数

	for i := 1; i <= input_count; i++ {
		fmt.Print("请输入数字:")
		fmt.Scanln(&input_num)

		if input_num > result {
			fmt.Printf("猜大了,还有%d次机会.\n", input_count-i)
		} else if input_num < result {
			fmt.Printf("猜小了,还有%d次机会.\n", input_count-i)
		} else if input_num == result {
			fmt.Println("猜对了.")
			break
		}
	}
	fmt.Printf("本次随机数为:%d\n", result)
}
