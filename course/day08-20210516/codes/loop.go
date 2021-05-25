package main

import (
	"fmt"
	"sync"
	"time"
)

// 闭包陷阱
func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(i int) {
			prefix := fmt.Sprintf("c%d", i+1)
			//
			for c := 'A'; c <= 'A'; c++ {
				fmt.Printf("%s: %c\n", prefix, c)
				time.Sleep(time.Millisecond)
			}
			wg.Done()
		}(i)
	}
	// time.Sleep(time.Microsecond)
	fmt.Println("wait")
	wg.Wait()
	fmt.Println("over")
}
