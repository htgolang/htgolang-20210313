package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		target int
		guess  int
		count  int
		flag   int
	)
again:
	count = 5
	flag = 0
	rand.Seed(time.Now().Unix()) // 以当前时间戳为种子进行随机
	target = rand.Intn(30)
	for count > 0 {
		fmt.Print("please input a number: ")
		fmt.Scanln(&guess)
		if guess > target {
			count--
			fmt.Printf("sry, bigger than target, and you have %d times\n", count)
			continue
		} else if guess < target {
			count--
			fmt.Printf("sry, smaller than target, and you have %d times\n", count)
			continue
		}
		fmt.Print("you win, again? (1:YES,0:NO)")
		fmt.Scanln(&flag)
		if flag == 1 {
			goto again
		} else {
			break
		}
	}
}
