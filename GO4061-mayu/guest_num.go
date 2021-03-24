package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	var Guess_num int 
	var juder string 
	END:
	rand.Seed(time.Now().Unix())
	ok_num := rand.Intn(100)
	fmt.Println(ok_num)
	for i:=0;i<5;i++{
 	fmt.Print("请输入查询的数字：")
	fmt.Scan(&Guess_num)
	if Guess_num == ok_num {
		fmt.Println("正确")
		fmt.Print("猜对了，是否继续 y/n: ")
		fmt.Scan(&juder)
		if juder == "y"{
			goto END 
		}else {
			return
		}
	}  else if Guess_num > ok_num{
		fmt.Println("猜大了")
	} else if Guess_num < ok_num{
		fmt.Println("猜小了")
	}
}
}
