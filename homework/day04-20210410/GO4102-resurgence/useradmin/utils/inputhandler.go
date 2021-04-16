package utils

import (
	"fmt"
)

func Input(prompt string) (text string) {
	fmt.Print(prompt)
	fmt.Scan(&text)
	return
}




