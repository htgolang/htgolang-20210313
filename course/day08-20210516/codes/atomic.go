package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&counter, -1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)

}
