package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		prefix := "c1"
		//
		for c := 'A'; c <= 'Z'; c++ {
			fmt.Printf("%s: %c\n", prefix, c)
			time.Sleep(time.Millisecond)
		}
		wg.Done()
	}()

	go func() {
		prefix := "c2"
		for c := 'A'; c <= 'Z'; c++ {
			fmt.Printf("%s: %c\n", prefix, c)
			time.Sleep(time.Millisecond)
		}
		wg.Done()
	}()

	fmt.Println("wait")
	wg.Wait()
	fmt.Println("over")
}
