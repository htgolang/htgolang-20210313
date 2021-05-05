package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//Check the parameters
	if len(os.Args) != 2 {
		log.Fatal("usage: count_lines.go [target file]")
	}
	//Provide the target file from the command line
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	fileLines := bufio.NewReader(file)
	//Start to count
	counts := 0
	for {
		_, err := fileLines.ReadString('\n')
		if err != nil {
			break
		}
		counts++
	}
	fmt.Println(counts)
}
