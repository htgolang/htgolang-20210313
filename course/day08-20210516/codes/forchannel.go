package main

import "fmt"

func main() {
	channel := make(chan int, 3)

	channel <- 1
	channel <- 2
	channel <- 3
	close(channel)

	// for v := range channel {
	// 	fmt.Println(v)
	// }
	for {
		if v, ok := <-channel; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}
