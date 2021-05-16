package main

import (
	"fmt"
	"time"
)

func After(d time.Duration) <-chan int {
	channel := make(chan int, 1)
	go func() {
		time.Sleep(d)
		channel <- 0
	}()
	return channel
}

func main() {
	fmt.Println(time.Now())
	<-After(1 * time.Second)
	fmt.Println(time.Now())

	fmt.Println(<-time.After(1 * time.Second)) // 延迟
	fmt.Println(time.Now())

	timeout := time.After(time.Second)
	fmt.Println(<-timeout)
	//fmt.Println(<-timeout)

	t := After(time.Second)
	<-t
	<-t
}
