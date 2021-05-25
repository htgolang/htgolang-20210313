package main

import (
	"fmt"
	"time"
)

func main() {
	// 多个管道进行读, 某一个管道读取成功就执行对应逻辑
	// 多个管道进行写, 某一个管道写成功就执行对应逻辑
	// 多个管道 某写进行读,某些进行写,某个管道 读/写成功执行对应逻辑
	// select case
	/*
		select {
		case value, ok := <-channel:
		case channel <- value:
		default:
		}*/

	channelV1 := make(chan int, 0)
	channelV2 := make(chan int, 0)

	go func() {
		time.Sleep(3 * time.Second)
		channelV2 <- 2
	}()

	fmt.Println("wait", time.Now())
	select {
	case v, ok := <-channelV1:
		fmt.Println("channelV1:", time.Now(), v, ok)
	case v, ok := <-channelV2:
		fmt.Println("channelV2:", time.Now(), v, ok)
	default:
		fmt.Println("default")
	}
	fmt.Println("over")
}
