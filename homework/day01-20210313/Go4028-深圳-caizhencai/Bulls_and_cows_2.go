package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
START:
	rand.Seed(time.Now().Unix())
	var guess_code int = rand.Int()
	// guess_code = 50
	var user_guess int
	var user_continue string = ""
	var guess_count int = 0
	for {
		if guess_count == 5 {
			fmt.Println("你已经连续5次没有猜中，自动退出猜数游戏。")
			break
		} else {
			guess_count++
		}
		fmt.Printf("你有5次猜数的机会，这是你第%d次猜数，请输入你猜测数字：", guess_count)
		fmt.Scan(&user_guess)
		if guess_code == user_guess {
			fmt.Printf("恭喜你猜对了，随机数为:%d \n", guess_code)
			fmt.Printf("是否继续猜数游戏？输入(y/n)选择:")
			fmt.Scan(&user_continue)
			if user_continue == "y" {
				goto START
			} else {
				break
			}

		} else if user_guess > guess_code {
			fmt.Printf("你猜的数字：%d,偏大了，请往小猜\n", user_guess)
		} else {
			fmt.Printf("你猜的数字：%d,偏小了，请往大猜。\n", user_guess)
		}

	}

}
