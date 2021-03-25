package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("C:\Users\Administrator\go\ihaveadream.txt")
	if err != nil {
		fmt.Println("Reading error", err)
	}
	fmt.Println("Contents of file:", string(data))
}
