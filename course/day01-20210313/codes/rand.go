package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	// 种子只设置一次

	fmt.Println(rand.Int())

	// 0 - 100
	fmt.Println(rand.Int() % 101) // 0-100
	fmt.Println(rand.Intn(101))

	// fmt.Println(rand.Intn(1))
}
