package main

import "fmt"

func main() {
	zero := make(chan int)
	one := make(chan int)
	// 5 => 0 5 => 1
	go func() {
		for i := 0; i < 6; i++ {
			zero <- 0
		}
	}()
	go func() {
		for i := 0; i < 4; i++ {
			one <- 1
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case <-one:
			fmt.Print(1)
		case <-zero:
			fmt.Print(0)
		}
	}

}
