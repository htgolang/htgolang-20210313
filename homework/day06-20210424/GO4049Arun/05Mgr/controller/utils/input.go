package utils

import "fmt"

func Input(prompt string) string {
	fmt.Print(prompt)
	var txt string
	fmt.Scan(&txt)
	return txt
}
