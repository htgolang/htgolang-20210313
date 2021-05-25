package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker := time.Tick(time.Second) // 每间隔N duration再管道中写入一次
	// fmt.Println(<-ticker)
	// fmt.Println(<-ticker)
	// fmt.Println(<-ticker)
	// for now := range ticker {
	// 	fmt.Println(now)
	// }

	ticker := time.NewTicker(time.Second)
	stop := make(chan int, 1)
	go func() {
		<-time.After(10 * time.Second)
		close(stop)
	}()

END:
	for {
		select {
		case <-stop:
			break END
		case now := <-ticker.C:
			fmt.Println(now)
		}
	}
	ticker.Stop()
}
