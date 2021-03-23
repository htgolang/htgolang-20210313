package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var (
		rndNum int = rand.Intn(100) //定义随机数
		usrNum int                  //定义用户输入数据
	)

	for i := 1; i <= 5; i++ {
		fmt.Printf("请输入数字：")
		fmt.Scan(&usrNum)

		if i > 6 {
			fmt.Println("超出5次，已退出")
			break
		}

		if usrNum > rndNum {
			fmt.Println("猜大了，请重新输入")
		} else if usrNum < rndNum {
			fmt.Println("猜小了，请重新输入")
		} else if usrNum == rndNum {
			fmt.Println("恭喜你猜对了！")
			break
		}

	}
	fmt.Println("随机数为：", rndNum)

}
