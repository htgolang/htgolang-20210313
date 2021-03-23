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
		contents[string(v)]++
	}
	//	fmt.Printf("%T", contents)
	fmt.Println(contents)
}
