package main

import (
	"fmt"
	"sync"
	"time"
)

// 要获取网站的图片
// 1.jpg, 2.jpg, ...., N.jpg
// 下载 m.jpg n.jpg

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

	// 生产者
	// 生成下载图片的例程
	go func(channel chan<- File) {
		for i := m; i <= n; i++ {
			channel <- File{fmt.Sprintf("%d.jpg", i), fmt.Sprintf("downlowd/%d.jpg", i)}
		}
		close(channel)
		wg.Done()
	}(channel)

	// 消费者
	for i := 0; i < N; i++ {
		// N个下载图片例程
		wg.Add(1)
		go func(channel <-chan File) {
			for f := range channel {
				download(f.name, f.path)
			}
			wg.Done()
		}(channel)
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
