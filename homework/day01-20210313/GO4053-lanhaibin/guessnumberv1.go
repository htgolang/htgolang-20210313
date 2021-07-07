package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 猜数字
// 只有5次机会，猜错则游戏结束。

func main() {
	var answer, guess, maxtry int
	rand.Seed(time.Now().UnixNano())
	answer = rand.Int() % 101
	maxtry = 5
	for i := 0; i < maxtry; i++ {
		fmt.Print("请输入一个0~100的数字:")
		fmt.Scanf("%d", &guess)
		if guess == answer {
			fmt.Println("恭喜你，猜对了!")
			return
		} else if guess > answer {
			fmt.Println("你猜的数字大了.")
		} else {
			fmt.Println("你猜的数字小了.")
		}
	}
	fmt.Println("游戏结束.")
}
