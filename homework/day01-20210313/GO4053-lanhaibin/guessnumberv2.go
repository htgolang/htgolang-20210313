package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	var answer, guess, maxtry int
	rand.Seed(time.Now().UnixNano())
	maxtry = 5
	var ask string
	answer = rand.Int() % 101
	for {
		for i:=0;i<maxtry;i++ {
			fmt.Println("请输入一个0~100的数:")
			fmt.Scanf("%d", &guess)
			if guess == answer {
				fmt.Println("恭喜你，猜对了!")
			} else if guess > answer {
				fmt.Println("你猜的数字大了.")
			} else {
				fmt.Println("你猜的数字小了.")
			}
		}
		if guess == answer {
			fmt.Print("是否继续游戏?(Y/N)")
		    fmt.Scanf("%s", &ask)
			if ask == "Y" || ask == "y" {
				answer = rand.Int() % 101
				continue
			}
		}
	fmt.Println("游戏结束!")
	break
	}
}
