package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/princebot/getpass"
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

func InputPassword(prompt string) string {
	password, err := getpass.Get(prompt)
	if err != nil {
		log.Fatal(err)
	}
	return string(password)
}
