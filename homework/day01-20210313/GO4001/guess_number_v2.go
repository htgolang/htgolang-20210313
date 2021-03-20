package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成一个随机数让用户猜，提示猜大猜小，最多5次哦
func main() {

START:
	var i int //i定义次数

	//random生成随机数
	rand.Seed(time.Now().UnixNano())
	var target = rand.Intn(101)
	for i = 1; i <= 5; i++ {
		var input int
		fmt.Printf("请输入一个1-100的数字，当前是第%d次猜:", i)
		fmt.Scan(&input)
		switch {
		case input < target:
			fmt.Printf("猜小了,你还有%d次机会\n", 5-i)
		case input > target:
			fmt.Printf("猜大了,你还有%d次机会\n", 5-i)
		case input == target:
			fmt.Printf("猜对了,你猜了%d次\n", i)
		default:
			break
		}

	}
	fmt.Printf("你没机会了，正确答案是%d\n", target)
	fmt.Printf("请问是否继续？(y/n)")
	var choice string
	fmt.Scan(&choice)
	if choice == "y" {
		goto START
	} else {
		fmt.Println("Bye-Bye")
	}
}
