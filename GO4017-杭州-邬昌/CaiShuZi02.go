package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	guess()

}


func guess(){

	rand.Seed(time.Now().Unix())
	var randNum = rand.Intn(100)
	fmt.Println(randNum)
	flag := 1

	for {

		var sal int
		var ifContinue string

		fmt.Println("请输入数字：")
		fmt.Scanln(&sal)
		if sal > randNum {
			fmt.Println("第",flag,"次猜，猜的大了")

		} else if sal < randNum {
			fmt.Println("第",flag,"次猜，猜的小了")
		} else if sal == randNum {
			fmt.Println("第",flag,"次猜，猜对了,继续吗?(y|n)")
			fmt.Scanln(&ifContinue)
			if ifContinue =="y"{
				randNum = rand.Intn(100)
				fmt.Println(randNum)
			} else {
				fmt.Println("游戏结束，请下次再来")
				break
			}

		}

		flag = flag + 1
		if flag > 5 {
			fmt.Println("猜了",flag-1,"次了不能再猜了")
			break
		}

	}

}