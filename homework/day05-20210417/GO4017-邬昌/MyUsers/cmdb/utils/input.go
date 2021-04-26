package utils

import "fmt"

func Input(str string) string {

	var text string
	fmt.Println(str)
	fmt.Scan(&text)
	return text

}
