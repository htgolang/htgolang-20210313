package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("bufio.txt")
	defer file.Close()
	// writer := bufio.NewWriter(file)
	writer := bufio.NewWriterSize(file, 16)
	fmt.Println(writer.WriteString("abcxyz123123123213213231"))
	writer.Flush() // 不要忘了flush

}
