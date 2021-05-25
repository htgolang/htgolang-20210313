package main

import (
	"fmt"
	"sync"
	"time"
)

// 要获取网站的图片
// 1.jpg, 2.jpg, ...., N.jpg
// 下载 m.jpg n.jpg

func download(wg *sync.WaitGroup, name, path string) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("download %s to %s\n", name, path)
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	m, n := 0, 100
	for i := m; i <= n; i++ {
		wg.Add(1)
		go download(&wg, fmt.Sprintf("%d.jpg", i), fmt.Sprintf("downlowd/%d.jpg", i))
	}

	// 用固定N个例程去请求 10 =>

	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}
