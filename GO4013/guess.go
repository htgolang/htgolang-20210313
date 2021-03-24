package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var in int
	rand.Seed(time.Now().UnixNano())
	shu := rand.Int() % 100
	for {
		fmt.Print("输入你猜的数字（0-100之间）：")
		fmt.Scanln(&in)
		if in > shu {
			fmt.Println("猜大了")
		} else if in < shu {
			fmt.Println("猜小了")
		} else {
			fmt.Println("猜对了")
			break
		}
	}

}
