package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var target = rand.Intn(100)
	var guessNum int
	var c string

	fmt.Print("本游戏只能猜五次，请输入你要猜的数字：")
	for i := 1; i <= 5; i++ {
		fmt.Scan(&guessNum)
		fmt.Printf("这是你第%d猜测。", i)
		if guessNum < target {
			fmt.Printf("你猜的数字太小了，你还有%d次机会。请继续输入你要猜测的数字：", 5-i)
		} else if guessNum > target {
			fmt.Printf("你猜的数字太大了，你还有%d次机会。请继续输入你要猜测的数字：", 5-i)
		} else if i == 5 && guessNum == target {
			fmt.Printf("好险啊，你终于猜对了，你用了%d次机会。请问是否继续游戏【Y/N】？", i)
			fmt.Scan(&c)
			if c == "Y" || c == "y" {
				i = 0
				fmt.Print("欢迎你再一次来到我们的游戏世界，请输入你要猜的数字：")
			} else {
				fmt.Printf("欢迎下次再来！！！")
				break
			}
		} else if i == 5 && guessNum != target {
			fmt.Printf("很抱歉，你耗尽了一生的气运也没能猜对，游戏结束。")
		} else {
			fmt.Printf("恭喜你才对，你只用%d次机会，太厉害了！！！请问是否继续游戏【Y/N】？\n", i)
			fmt.Scan(&c)
			if c == "Y" || c == "y" {
				i = 0
				fmt.Print("欢迎你再一次来到我们的游戏世界，请输入你要猜的数字：")
			} else {
				fmt.Printf("欢迎下次再来！！！")
				break
			}
		}
	}
}
