package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		go func() {
			for {
			}
		}()
		for {
		}

	}()

	time.Sleep(3 * time.Second)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOROOT())
}
