package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
START:
	var (
		input_num    int
		input_choice string
	)

	rand.Seed(time.Now().UnixNano()) //设置随机数种子,种子只需要设置一次
	computer_num := rand.Intn(101)
	fmt.Println("猜数字游戏0-100，你有五次机会")
	for i := 5; i > 0; i-- {
		fmt.Println(computer_num)
		fmt.Print("请输入你的数字：")
		fmt.Scan(&input_num)
		if input_num == computer_num {
			fmt.Printf("恭喜你，猜对了\n结果就是%d，请问是否还继续？[y/n]: ", input_num)
			fmt.Scan(&input_choice)
			if input_choice == "y" {
				goto START
			} else {
				fmt.Println("感谢参与")
			}
		} else if i == 1 {
			fmt.Printf("你已经猜错了5次，现在退出！")
			break
		} else if input_num > computer_num {
			fmt.Printf("你猜错了，再往小一点猜\n你还剩余%d次机会", i-1)
		} else {
			fmt.Printf("你猜错了，再大一点猜\n你还剩余%d次机会", i-1)
		}

	}

}
