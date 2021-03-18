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
	)
	count = 5
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
		fmt.Print("you win")
		break
	}
}
