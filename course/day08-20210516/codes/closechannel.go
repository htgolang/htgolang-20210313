package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan int)
	// bufferChannel := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		// 关闭管道
		time.Sleep(3 * time.Second)
		v, ok := <-channel

		fmt.Println(v, ok)
		v, ok = <-channel
		fmt.Println(v, ok)
		wg.Done()
	}()

	go func() {
		// 不能针对关闭的管道写数据
		// 针对已经关闭的管道可以进行读取
		channel <- 1
		wg.Done()
		// close(channel)
	}()

	wg.Wait()
}
