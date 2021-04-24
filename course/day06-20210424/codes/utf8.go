package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("gbk.txt")
	reader := bufio.NewReader(file)
	fmt.Println(reader.ReadString('\n'))
}
