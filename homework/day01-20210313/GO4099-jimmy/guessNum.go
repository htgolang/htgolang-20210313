package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX int = 5

func main() {
	var res int
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
			break
		}
	}
}
