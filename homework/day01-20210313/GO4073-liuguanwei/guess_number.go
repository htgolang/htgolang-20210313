package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
start:
	//Define the random number.
	count := 0
	rand.Seed(time.Now().UnixNano())
	generate_number := rand.Intn(100)
	var input_number int
middle:
	//Exit the program if five times.
	count += 1
	if count == 6 {
		fmt.Println("Already five times.")
		os.Exit(3)
	}
	fmt.Println("Pls enter a number(0-100): ")
	fmt.Scan(&input_number)
	//Compare the number in for loop.
	for {
		//Bigger and go back to guess again.
		if generate_number > input_number {
			fmt.Println("Pls enter a bigger one.")
			goto middle
			//Smaller and go back to guess again.
		} else if generate_number < input_number {
			fmt.Println("Pls enter a smaller one.")
			goto middle
			//Correct
		} else {
			fmt.Println("Bingo")
			//Guess again if yes and exit if no
			fmt.Println("Would you like to try again?[yes|no]: ")
			var confirm_try string
			fmt.Scan(&confirm_try)
			if confirm_try == "yes" {
				goto start
			} else {
				os.Exit(3)
			}
		}
		break
	}
}
