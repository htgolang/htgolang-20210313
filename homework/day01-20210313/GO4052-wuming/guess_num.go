package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//生成一个50以内的随机数
	rand.Seed(time.Now().UnixNano())
	var tmp int
	var target int
	tmp = rand.Int()
	target = tmp % 51
	println(target)
	//用户输入一个数字
	var num int

	//判断大小
	for {
		println("请输入数字：")
		fmt.Scan(&num)
		if num > target {
			println("太大了")
		} else if num < target {
			println("太小了")
		} else {
			println("猜对了")
			break
		}
	}

}
