package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //随机种子
	num2 := rand.Intn(10)
	var num1 int
	for i := 1; i <= 5; i++ {
		fmt.Println("Please enter the number you guessed（1-10）：")
		fmt.Scan(&num1)
		if num1 == num2 {
			fmt.Println("Congratulations, you got it right!")
			break
		} else if num1 < num2 {
			fmt.Printf("Guess it's small! %d chances left\n",5-i)
		} else if num1 > num2 {
			fmt.Printf("Guess it's big! %d chances left\n",5-i)
		}

	}
}