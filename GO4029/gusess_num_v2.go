package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var input_num int
	var tag string
	fmt.Printf("猜数字游戏开始。。。")
	rand.Seed(time.Now().UnixNano())
	rand_num := rand.Intn(50)
	for i := 1; i <= 5; i++ {
		fmt.Printf("请输入数字：")
		fmt.Scan(&input_num)
		if input_num > rand_num {
			fmt.Printf("猜大了\n")
			fmt.Printf("是否继续?y/n")
			fmt.Scan(&tag)
			if tag == "n" {
				break
			} else {
				continue
			}
		} else if input_num < rand_num {
			fmt.Printf("猜小了\n")
			fmt.Printf("是否继续?y/n")
			fmt.Scan(&tag)
			if tag == "n" {
				break
			} else {
				continue
			}
		} else {
			fmt.Printf("猜对了\n")
			fmt.Printf("是否继续?y/n")
			fmt.Scan(&tag)
			if tag == "n" {
				break
			} else {
				continue
			}
		}
	}
}
