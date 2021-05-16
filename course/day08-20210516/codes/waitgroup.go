package main

import (
	"fmt"
	"sync"
	"time"
)

// 打印A->Z
func printChars(wg *sync.WaitGroup, prefix string) {
	// stage B
	defer wg.Done()
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Millisecond)
	}
	// stage C

	wg.Add(1)
	go childChars(wg, prefix)
}

func childChars(wg *sync.WaitGroup, prefix string) {
	defer wg.Done()

	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("child.%s: %c\n", prefix, c)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup // 定义计数信号量
	/*
		stage A
		go func(){
			stage B
			...
			有可能报错
			...
			stage C
		}
		stage D
	*/
	// wg.Add(N) => 启动例程之前执行 stage A, 再计数信号量中+N
	// wg.Done() => 例程执行结束后调用(例程函数退出时) stage C, 如何保证Done一定执行 stage B+defer, 当函数执行结束对计数信号量减1
	// wg.Wait() => 启动所有例程后调用, 等待计数为0的之后执行(不是0的时候一致等待)

	// stage A
	wg.Add(1)
	go printChars(&wg, "c1")
	wg.Add(1)
	go printChars(&wg, "c2")
	// stage D
	fmt.Println("wait")
	wg.Wait()
	fmt.Println("over")

	// waitgroup.Add数量 <= 例程数量=>Done()
	// 某些例程不需要等待结束 即可结束主例程(程序)
}
