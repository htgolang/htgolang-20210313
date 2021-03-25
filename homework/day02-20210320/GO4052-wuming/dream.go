package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

func main() {

	f, err := ioutil.ReadFile("Ihaveadream.txt")
	if err != nil {
		fmt.Println("read fail", err)
	}

	toCount := make(map[string]int)

	for _, v := range string(f) {
		if unicode.IsLetter(v){
			toCount[string(v)]++
		}else{
			continue
		}

		for i,v := range toCount{
			fmt.Println(i,v)
		}
	}


}