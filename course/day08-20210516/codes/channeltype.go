package main

import "fmt"

func main() {
	// 读管道
	// 写管道

	// 再某个函数中只需要读,或者只需要写的时候
	// 为防止再只读函数中误写, 再只写函数中误读
	// 可以将管道声明为只读,只写
	channel := make(chan int, 10)
	var readChannel <-chan int
	var writeChannel chan<- int

	readChannel = channel
	writeChannel = channel

	writeChannel <- 1
	writeChannel <- 2

	fmt.Println(<-readChannel)
	fmt.Println(<-readChannel)

	/*
		var ch chan<- int = make(chan int, 10)
		ch <- 1

		fmt.Println(<-ch)
	*/
}
