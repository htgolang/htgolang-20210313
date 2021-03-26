package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	c := string(content)

	count := map[string]int{}

	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, i := range c {
		if strings.Contains(characters, string(i)) {
			// fmt.Printf("%T, %v", i, i)
			count[string(i)]++
		}
		// else {
		// 	fmt.Printf("%T, %v\n", string(i), string(i))
		// }
		// fmt.Printf("%T, %v\n", i, i)
		// fmt.Println(string(i))
	}
	// fmt.Println(count)
	for k, v := range count {
		fmt.Printf("字符: %s, 出现次数: %v\n", k, v)
	}
}
