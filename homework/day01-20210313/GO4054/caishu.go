package main

import (
	"fmt"
	"math/rand"
	"time"
)
var guess int
func main(){
	rand.Seed(time.Now().Unix())
	var target int = rand.Intn(50)
	for{
		fmt.Println("请输入猜测数字")
		fmt.Scan(&guess)
		if guess > target{
			fmt.Println("太大了")
		} else if guess < target{
			fmt.Println("太小了")
		} else{
			fmt.Println("猜对了")
		}
	}

}