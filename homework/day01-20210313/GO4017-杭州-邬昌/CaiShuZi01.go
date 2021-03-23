package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	rand.Seed(time.Now().Unix())
	var randNum = rand.Intn(100)
	flag := 1

	for {

		var sal int

		fmt.Println("请输入数字：")
		fmt.Scanln(&sal)
		if sal > randNum {
			fmt.Println("第",flag,"次猜，猜的大了")

		} else if sal < randNum {
			fmt.Println("第",flag,"次猜，猜的小了")
		} else if sal == randNum {
			fmt.Println("第",flag,"次猜，猜对了，这个数字是",sal)
			break

		}

		flag = flag + 1
		if flag > 5 {
			fmt.Println("猜了5次了不能再猜了")
			break
		}

	}

}