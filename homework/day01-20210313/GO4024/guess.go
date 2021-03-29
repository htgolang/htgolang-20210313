package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Int())

	//随机数在0-100之间
	randNum := rand.Int() % 100
	//fmt.Println(randNum)

	//定义变量接收输入数字
	var guessNum int
	
	//最多猜5次
	for i := 5; i > 0; i --{
		//接收猜的数字存到变量guessNum
		fmt.Println("请输入你猜的数字(0-100): ")
		fmt.Scan(&guessNum)
		//fmt.Println(guessNum)
		//判断猜的数字与随机数的大小
		if guessNum > randNum {
			fmt.Println("猜大了")
		} else if guessNum < randNum {
			fmt.Println("猜小了")
		} else {
			//猜对退出循环
			fmt.Println("太棒了，猜对了")
			break
		}
		//每次猜完提示还剩余次数
		fmt.Println("你还有", i-1, "次机会")
	}
}