package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	guessNumber()
}


func guessNumber() {
	rand.Seed(time.Now().UnixNano()) //随机种子
	num2 := rand.Intn(10)   //10以内的随机数
	var num1 int           //接收用户输入的变量
	var isyes string      //是否继续的变量
	for i := 1; i <= 5; i++ {
		fmt.Print("Please enter the number you guessed（1-10）：")
		fmt.Scan(&num1)
		if num1 == num2 {			 
			fmt.Println("Congratulations, you got it right!")
			START:    
			fmt.Print("Whether to continue(y or n): ")
			fmt.Scan(&isyes)
			if isyes == "y" {
				guessNumber()
			} else if isyes == "n"{
				fmt.Println("Game over.")
				goto END
			} else {
				fmt.Println("input error, again.")
				goto START
			}
		} else if num1 < num2 {
			fmt.Printf("Guess it's small! %d chances left\n",5-i)
		} else if num1 > num2 {
			fmt.Printf("Guess it's big! %d chances left\n",5-i)
		}	
	}
	END:
}