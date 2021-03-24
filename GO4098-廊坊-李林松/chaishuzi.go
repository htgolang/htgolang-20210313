package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessNum() {
	var target int
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(11)
	for {
		fmt.Print("请输入十以内的数字:")
		fmt.Scan(&target)
		if target < randNum {
			fmt.Println("输入的数字太小请重新输入")
		} else if target > randNum {
			fmt.Println("输入的数字太大请重新输入")
		} else {
			fmt.Printf("猜对了, 数字是: %d\n", randNum)
			break
		}
	}

}
func main() {
	guessNum()
}
