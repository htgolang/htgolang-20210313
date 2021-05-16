package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 事件广播
	// ctx, cannel := context.WithCancel(ctx)
	// cannel() => 广播消息
	// ctx

	channel := make(chan int, 1)

	go func() {
		i := 0
		for {
			channel <- i
			time.Sleep(time.Second * 2)
			i += 1
		}
	}()

	// ctx, cannel := context.WithCancel(context.Background())
	// go func() {
	// 	time.Sleep(time.Second * 10)
	// 	fmt.Println("cannel")
	// 	cannel()
	// }()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// cannel()

END:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("interrupt")
			break END
		case v, ok := <-channel:
			if !ok {
				break END
			} else {
				fmt.Println(v)
			}
		}
	}

}
