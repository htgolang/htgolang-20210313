package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 5 //guess counts
	var inputNum int
	//setting the seed of random, only one time setting is enough
	rand.Seed(time.Now().UnixNano())
	var num int = 100
	var guessNum int = rand.Intn(num) + 1 //range in [1,num]
	for i := 1; i <= count; i++ {
		fmt.Printf("Please input a number that ranges from 1 to %v\n", num)
		fmt.Scan(&inputNum)
		if inputNum == guessNum {
			fmt.Printf("Congratulation, you are right, which num is %v\n", guessNum)
		} else if inputNum > guessNum {
			fmt.Printf("The number inputted is bigger than the right number! Please input again!\n")
		} else if inputNum < guessNum {
			fmt.Printf("The number inputted is smaller than the right number! Please input again!\n")
		}
	}
}

/*
> go run .\guessNum.go
Please input a number that ranges from 1 to 100
50
The number inputted is bigger than the right number! Please input again!
Please input a number that ranges from 1 to 100
25
The number inputted is bigger than the right number! Please input again!
Please input a number that ranges from 1 to 100
12
The number inputted is bigger than the right number! Please input again!
Please input a number that ranges from 1 to 100
5
Congratulation, you are right, which num is 5
*/
