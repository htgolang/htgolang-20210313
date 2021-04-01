package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const MAX int = 5

func main() {
	var res int
	var again string

START:
	rand.Seed(time.Now().Unix())
	random := rand.Intn(100)

	for i := 1; i <= MAX; i++ {
		fmt.Print("请输入一个整数: ")
		if _, err := fmt.Scanln(&res); err != nil {
			fmt.Println("请重新输入")
			continue
		}

		if res > random {
			fmt.Printf("猜大了，您还有%d次机会\n", MAX-i)
		} else if res < random {
			fmt.Printf("猜小了，您还有%d次机会\n", MAX-i)
		} else {
			fmt.Printf("猜对了，您用了%d次机会\n", i)
		AGAIN:
			fmt.Print("是否继续(y/n): ")
			fmt.Scanln(&again)
			switch again {
			case "y", "Y":
				goto START
			case "n", "N":
				os.Exit(0)
			default:
				fmt.Println("输入错误")
				goto AGAIN
			}
		}
	}
}
