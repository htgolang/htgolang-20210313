package main

import (
	"fmt"
	"strconv"
)

func main() {
	// "123" => int

	if i, err := strconv.Atoi("10"); err == nil {
		fmt.Printf("%T %v\n", i, i)
	} else {
		fmt.Println("error:", err)
	}
	// int => string
	fmt.Printf("%T %v\n", strconv.Itoa(10), strconv.Itoa(10))
	// "12.2" => float64
	f, err := strconv.ParseFloat("12.2", 64)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("%T, %v\n", f, f)
	}
	// float64 => string
	fmt.Println(strconv.FormatFloat(12.2, 'E', -1, 64))
	fmt.Println(strconv.FormatFloat(12.2, 'f', -1, 64))

	// string => int8 int16 int32 int64
	// ParseInt ParseUint
	// int64, int32 int16 int8 => string
	// FormatInt FormatUint
}
