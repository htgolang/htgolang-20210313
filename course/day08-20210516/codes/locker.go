package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int

	var wg sync.WaitGroup

	var locker sync.Mutex
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				locker.Lock()
				counter += 1
				locker.Unlock()
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < 10000; j++ {
				locker.Lock()
				counter -= 1
				locker.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)

}
