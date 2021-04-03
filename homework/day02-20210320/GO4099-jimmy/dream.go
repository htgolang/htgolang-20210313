package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("dream.txt")
	if err != nil {
		panic(err)
	}

	res := make(map[string]int)
	for _, d := range data {
		if d >= 'a' && d <= 'z' || d >= 'A' && d <= 'Z' {
			res[string(d)]++
		}
	}

	for k, v := range res {
		fmt.Println(k, v)
	}
}
