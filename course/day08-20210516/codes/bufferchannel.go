package main

import (
	"fmt"
	"time"
)

func main() {
	// var channel chan string = make(chan string, 3)
	// var channel = make(chan string, 3)
	channel := make(chan string, 3)

	fmt.Println(len(channel))
	channel <- "a"
	fmt.Println(len(channel))
	channel <- "b"
	channel <- "c"

	fmt.Println(len(channel))
	fmt.Println(<-channel) // a

	fmt.Println(len(channel))

	channel <- "d"
	fmt.Println(<-channel) // b
	fmt.Println(<-channel) // c
	fmt.Println(<-channel) // d

	go func() {
		time.Sleep(5 * time.Second)
		channel <- "x"
	}()

	fmt.Println("wait")
	fmt.Println(<-channel) // x
	fmt.Println("over")
}
