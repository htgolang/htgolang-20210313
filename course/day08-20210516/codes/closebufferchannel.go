package main

import "fmt"

func main() {
	bufferChannel := make(chan int, 3)

	bufferChannel <- 1

	close(bufferChannel)
	v, ok := <-bufferChannel

	fmt.Println(v, ok)
	v, ok = <-bufferChannel

	fmt.Println(v, ok)
}
