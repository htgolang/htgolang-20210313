package main

import (
	"fmt"
	"math/rand"
	"os"
)

func GuessNumber(usernumber int) {
	random_number := rand.Intn(1)
	switch {
	case random_number < usernumber:
		fmt.Println("big")
	case random_number > usernumber:
		fmt.Println("small")
	case random_number == usernumber:
		var choose string
		fmt.Println("success")
		fmt.Println("Whether or not to continue？(yes/no)")
		fmt.Scan(&choose)
		if choose == "yes" {
			break
		} else if choose == "no" {
			os.Exit(1)
		}
	}
}

func main() {
	var input int
	counts := int(0)
	for counts <= 5 {
		fmt.Println("Please enter the number 1-100: ")
		fmt.Scan(&input)
		GuessNumber(input)
		counts++
		fmt.Printf("%d opportunities have been used \n", counts)
		fmt.Println()
		continue
	}
}
