package main

import (
	"fmt"
	"sync"
	"time"
)

type File struct {
	name string
	path string
}

func download(name, path string) {
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("download %s to %s\n", name, path)
}

func main() {
	start := time.Now()
	N := 10
	m, n := 1, 100
	channel := make(chan File, N)

	wg := sync.WaitGroup{}
	wg.Add(1)

	//生产者
	//生成下载图片的例程
	go func(channel chan<- File) { //只写
		for i := m; i <= n; i++ {
			channel <- File{fmt.Sprintf("%d.jpg", i), fmt.Sprintf("download/%d.jpg", i)}
		}
		close(channel)
		wg.Done()
	}(channel)

	//消费者
	//N个下载图片的例程
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(channel <-chan File) { //<---- 只读
			for f := range channel {
				download(f.name, f.path)
			}
			wg.Done()
		}(channel)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start)) //应该比较顺序编程慢，因为例程之间切换也消耗时间
}
