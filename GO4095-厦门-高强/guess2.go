package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var target = rand.Intn(100)
	var guess int
	const maxCounts int = 5

NEXT:
	fmt.Println("现在开始猜数游戏，数字范围为1--100")
	for i := 1; i <= maxCounts; i++ {
		fmt.Printf("请输入你猜测的数字：")
		_, _ = fmt.Scan(&guess)
		if guess < target {
			fmt.Println("你输入的数字比目标要小，请重新输入！")
		} else if guess > target {
			fmt.Println("你输入的数字比目标要大，请重新输入！")
		} else {
			fmt.Println("输入正确，是否还要继续 yes or no ?")
			var answer string
			_, _ = fmt.Scan(&answer)

			if answer == "yes" {
				goto NEXT
			}
			break
		}
	}
}
