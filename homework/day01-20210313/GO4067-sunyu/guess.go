package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var guess int
	var num int
	rand.Seed(time.Now().Unix())
	num = rand.Intn(100)
	var yn string
END:
	for t := 0; t < 5; t++ {
		fmt.Println("请猜一个数字：")
		fmt.Scan(&guess)
		if guess > num {
			fmt.Println("猜大了")
		} else if guess < num {
			fmt.Println("猜小了")

		} else {
			fmt.Println("Bingo!")
			fmt.Println("是否继续y/n?")
			fmt.Scan(&yn)
			if yn == "y" {
				goto END
			} else {
				break END
			}
		}
	}
	fmt.Printf("正确数字是%d", num)
	fmt.Println("是否继续y/n?")
	fmt.Scan(&yn)
	if yn == "y" {
		goto END
	}

}
