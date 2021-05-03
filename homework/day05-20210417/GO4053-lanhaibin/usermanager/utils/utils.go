package utils

import (
	"fmt"
	"strconv"
	"time"
)

func Input(prompt string) string {
	var s string
	fmt.Print(prompt)
	fmt.Scan(&s)
	return s
}

func InputInt(prompt string) int {
	s := Input(prompt)
	i, _ := strconv.Atoi(s)
	return i
}

func InputTime(prompt string) time.Time {
	t, _ := time.Parse("2006-01-02", Input(prompt))
	return t
}
