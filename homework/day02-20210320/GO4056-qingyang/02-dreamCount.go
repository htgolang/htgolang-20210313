package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkChar(v byte) bool {
	if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
		return true
	}
	return false
}

func main() {
	file, err := os.Open("dream")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	text, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	countMap := make(map[string]int)
	for _, v := range text {
		if checkChar(v) {
			countMap[string(v)]++
		}
	}
	fmt.Println(countMap)
}
