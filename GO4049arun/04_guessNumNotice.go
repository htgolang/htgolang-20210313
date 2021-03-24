package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 5 //guess counts
	var tag int = 0
	var inputNum int
	//setting the seed of random, only one time setting is enough
	rand.Seed(time.Now().UnixNano())
	var num int = 100
	var guessNum int = rand.Intn(num) + 1 //range in [1,num]
	var ifGo string
	for i := 1; i <= count; i++ {
		fmt.Printf("Please input a number that ranges from 1 to %v\n", num)
		fmt.Scan(&inputNum)
		tag++
		if inputNum == guessNum {
			fmt.Printf("Congratulation, you are right, which num is %v\n", guessNum)
		} else if inputNum > guessNum {
			if tag == 5 {
				fmt.Printf("您的次数已超过上限%v次,将退出程序", count)
				return
			}
			fmt.Printf("大了!是否继续y/n\n")
			fmt.Scan(&ifGo)
			if "y" == ifGo {
				continue
			} else {
				return
			}
		} else if inputNum < guessNum {
			if tag == 5 {
				fmt.Printf("您的次数已超过上限%v次,将退出程序", count)
				return
			}
			fmt.Printf("小了!是否继续y/n\n")
			fmt.Scan(&ifGo)
			if "y" == ifGo { //这里应该有数据类型判断,暂时忽略
				continue
			} else {
				return
			}
		}
	}
}

/*
> go run .\03_guessNumNotice.go
Please input a number that ranges from 1 to 100
50
小了!是否继续y/n
y
Please input a number that ranges from 1 to 100
76
大了!是否继续y/n
y
Please input a number that ranges from 1 to 100
69
小了!是否继续y/n
y
Please input a number that ranges from 1 to 100
71
小了!是否继续y/n
y
Please input a number that ranges from 1 to 100
72
Congratulation, you are right, which num is 72
-------------------------------------------------------------------------
> go run .\03_guessNumNotice.go
Please input a number that ranges from 1 to 100
99
大了!是否继续y/n
y
Please input a number that ranges from 1 to 100
98
大了!是否继续y/n
y
Please input a number that ranges from 1 to 100
97
大了!是否继续y/n
y
Please input a number that ranges from 1 to 100
96
大了!是否继续y/n
y
Please input a number that ranges from 1 to 100
95
您的次数已超过上限5次,将退出程序
*/
