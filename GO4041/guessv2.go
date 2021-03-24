package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机种子
	rand.Seed(time.Now().Unix())

	// 申明猜测变量int 继续猜测string
	var j int
	var k string

	// 生成随机数
a:
	r := rand.Intn(100)

	for i := 5; i >= 1; i-- {
		fmt.Printf("请输入您猜测的数字：还剩%d次！", i)
		fmt.Scanf("%d", &j)

		if j == r {
			fmt.Println("恭喜您猜对了！")
			fmt.Println("请问是否继续：y/n")
			fmt.Scanf("%s", &k)
			if k == "y" {
				goto a
			}
			break
		} else if j > r {
			fmt.Println("猜大了！")
		} else if j < r {
			fmt.Println("猜小了！")
		}
		if i == 1 {
			fmt.Printf("结果是：%d", r)
		}
	}
}
