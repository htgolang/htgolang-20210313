package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	txt, _ := ioutil.ReadFile("ihaveadream.txt")
	//	fmt.Println(string(txt))
	contents := make(map[string]int)
	for _, v := range string(txt) {
		if string(v) >= "a" && string(v) <= "z" || string(v) >= "A" && string(v) <= "Z" {
			contents[string(v)]++
		}
	}
	//	fmt.Printf("%T", contents)
	fmt.Println(contents)
}
