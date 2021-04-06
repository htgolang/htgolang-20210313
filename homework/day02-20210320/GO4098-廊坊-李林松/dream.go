package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func traverseArticle(path string) map[string]int {
	var statistics = make(map[string]int)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	txt, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(txt); i++ {
		if (txt[i] >= 65 && txt[i] <= 90) || (txt[i] >= 97 && txt[i] <= 122) {
			statistics[string(txt[i])]++
		}
	}
	return statistics
}

func main() {
	path := "dream.txt"
	ret := traverseArticle(path)
	for k, v := range ret {
		fmt.Println(k, v)
	}
}
