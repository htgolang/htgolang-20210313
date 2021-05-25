package main

import "fmt"

func main() {
	var channel chan int = make(chan int)

	for i := 0; i < 2; i++ {
		go func(i int) {
			for c := 'A'; c <= 'Z'; c++ {
				fmt.Printf("%d, %c\n", i, c)
			}

			// 执行完例程后写入数据
			channel <- 0
		}(i)
	}

	fmt.Println("wait")
	for i := 0; i < 2; i++ {
		<-channel
	}
	fmt.Println("over")
}
