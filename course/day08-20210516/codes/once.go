package main

import (
	"fmt"
	"sync"
)

func conf() {
	fmt.Println("conf")
}

func main() {
	var once sync.Once
	once.Do(conf)
	once.Do(conf)
	once.Do(conf)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(conf)
			wg.Done()
		}()
	}

	wg.Wait()
}
