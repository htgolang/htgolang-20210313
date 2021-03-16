// Author: kalidio
// Date: 2021-03-16
// Content: Day01 practice
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Multiplication is a function to calculate the multiplication formula.
func Multiplication() {
	fmt.Println("2、打印乘法口诀")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %-2d ", j, i, j*i)
		}
		fmt.Println()
	}
}

// PlayGGame is a function to guess num game.
func PlayGGame() {
	fmt.Println("3,4、猜数字游戏")
AGAIN:
	var GuessNum int
	var Answer string
	Times, MaxNum := 5, 10 // number of times, the maximum number
	// initialize random number
	rand.Seed(time.Now().UnixNano())
	GameNum := rand.Intn(MaxNum) + 1
	for { // loop
		fmt.Printf("Note: you have %d chances! ", Times)
		fmt.Printf("Please input your number(1 ~ %d): ", MaxNum)
	REINPUT:
		fmt.Scanf("%d\n", &GuessNum)
		// fmt.Printf("%d",GuessNum)
		if GuessNum > MaxNum || GuessNum <= 0 { // illegal input
			goto REINPUT
		} else if GuessNum < GameNum {
			fmt.Println("The number is too small!")
		} else if GuessNum > GameNum {
			fmt.Println("The number is too big!")
		} else {
			fmt.Println("You are so smart, you guessed it!")
			// break   // if not loop the game
		SELECT:
			fmt.Print("Do you want to play again?(y/n): ")
			// after guessing, can choose to continue the play
			fmt.Scanf("%s", &Answer)
			if Answer == "y" {
				goto AGAIN
			} else if Answer == "n" {
				// fmt.Println("The game is over.")
				break
			} else {
				// fmt.Println("You input is illegal, please re-input.")
				// fmt.Println()
				goto SELECT
			}
		}
		Times--
		if Times == 0 {
			fmt.Println("What a pity! you almost guessed it, please try again.")
			break
			// if want to re-play the game, the same above.
		}
	}
}

// YhTriangle is a function to print Yanghui Triangle.
func YhTriangle() {
	fmt.Println("5、打印杨辉三角形")
	const n int = 10
	var t [n][n]int
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ") // print spaces
		}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				fmt.Print(" ") // print spaces
				//     1      t[0][0] =
				//    1 1     t[1][0] = 1, t[1][1] = 1
				//   1 2 1    t[2][0] = 1, t[2][1] = 2, t[2][2] = 1
				t[i][j] = 1 // t[*][0] = t[i][j] = 1
			} else {
				//     1      t[0][0] =
				//    1 1     t[1][0] = 1, t[1][1] = 1
				//   1 2 1    t[2][0] = 1, t[2][1] = 2, t[2][2] = 1
				//  1 3 3 1   t[3][0] = 1, t[3][1] = 3, t[3][2] = 3, t[3][3] = 1
				fmt.Print(" ")                    // print spaces
				t[i][j] = t[i-1][j-1] + t[i-1][j] // t[3][2] = t[2][1] + t[2][2]
			}
			fmt.Printf("%d", t[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// 打印乘法口诀
	Multiplication()
	// 打印杨辉三角形
	YhTriangle()
	// 猜数字游戏
	PlayGGame()
}
