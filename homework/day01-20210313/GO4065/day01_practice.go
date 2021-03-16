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
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}

// PlayGGame is a function to guess num game.
func PlayGGame() {
	fmt.Println("3,4、猜数字游戏")
AGAIN:
	var GuessNum int
	var x string
	c, max := 5, 100 // number of times, the maximum number
	// initialize random number
	rand.Seed(time.Now().UnixNano())
	GameNum := rand.Intn(max) + 1
REINPUT:
	// loop
	for {
		fmt.Printf("Note: you have %d chances! ", c)
		fmt.Printf("Please input your number(0 ~ %d):\t", max)
		fmt.Scanf("%d\n", &GuessNum)
		if GuessNum > max || GuessNum < 0 {
			// illegal input
			fmt.Println("You input is illegal, please re-input.\t")
			goto REINPUT
		} else if GuessNum < GameNum {
			fmt.Println("The number is too small!")
		} else if GuessNum > GameNum {
			fmt.Println("The number is too big!")
		} else {
			fmt.Println("Congratulations, you guessed it!")
			// break   // do not loop the game
			fmt.Print("Do you want to play again?(y/n)\t")
		SELECT:
			// after guessing, can choose to continue the play
			fmt.Scanf("%s", &x)
			if x == "y" {
				// TODO：输入 y 后，loop 已执行了一次问题，如下：
				// Note: you have 5 chances! Please input your number:The number is too small!
				// Note: you have 4 chances! Please input your number:
				goto AGAIN
			} else if x == "n" {
				fmt.Println("The game is over.")
				break
			} else {
				fmt.Println("You input is illegal, please re-input.")
				fmt.Print("Please re-input your choice:\t")
				goto SELECT
			}
		}
		c--
		if c == 0 {
			fmt.Println("The game is over.")
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
