package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机种子
	rand.Seed(time.Now().Unix())

	// 申明猜测变量int
	var j int

	// 生成随机数
	r := rand.Intn(100)

	for i := 5; i >= 1; i-- {
		fmt.Printf("请输入您猜测的数字：还剩%d次！", i)
		fmt.Scanf("%d", &j)

		if j == r {
			fmt.Println("恭喜您猜对了！")
			break
		} else if j > r {
			fmt.Println("猜大了！")
		} else if j < r {
			fmt.Println("猜小了！")
		}
	}

	fmt.Printf("结果是：%d", r)
}
