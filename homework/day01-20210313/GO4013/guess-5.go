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
	cishu := 5
	for i := 1; i <= cishu; i++ {
		fmt.Print("输入你猜的数字（0-100之间）：")
		fmt.Scanln(&in)
		if in > shu {
			fmt.Printf("猜大了，机会还有%d次\n", cishu-i)
		} else if in < shu {
			fmt.Printf("猜小了，机会还有%d次\n", cishu-i)
		} else {
			fmt.Println("猜对了")
			break
		}
	}

}
