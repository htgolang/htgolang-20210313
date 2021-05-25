package main

import (
	"context"
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
	time.Sleep(3 * time.Second)
	// time.Sleep(3 * time.Millisecond)
	fmt.Printf("download %s to %s\n", name, path)
}

func main() {

	// 只下载10s秒, 10s内下载多少文件都OK

	// 10s计时器
	// close chan 10 后关闭管道 其他例程可以读取<-close // 广播消息
	// 关闭例程(生产图片/下载图片)
	// 如何关闭
	// 生产者
	/*
		select {
		case <-close:
			break
		case channel <- xxxx:
		}
	*/
	// 消费
	/*
		select {
		case <-close:
			break
		case f <-channel:
		}
	*/

	start := time.Now()
	N := 10
	m, n := 1, 100
	channel := make(chan File, N)

	// closeChannel := make(chan int, 1) //关闭管道
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	wg := sync.WaitGroup{}

	wg.Add(1)

	// 生产者
	// 生成下载图片的例程
	go func(channel chan<- File) {
	END:
		for i := m; i <= n; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("producer channel close")
				break END
			case channel <- File{fmt.Sprintf("%d.jpg", i), fmt.Sprintf("downlowd/%d.jpg", i)}:

			}

		}
		close(channel)
		wg.Done()
	}(channel)

	// 消费者
	for i := 0; i < N; i++ {
		// N个下载图片例程
		wg.Add(1)
		go func(channel <-chan File, i int) {
		END:
			for {
				select {
				case <-ctx.Done():
					fmt.Println("consumer channel close", i, "interrupt")
					break END
				case v, ok := <-channel:
					if !ok {
						fmt.Println("consumer channel close", i, "channel")
						break END
					} else {
						download(v.name, v.path)
					}
				}
			}
			wg.Done()
		}(channel, i)
	}

	// go func() {
	// 	// 只需要执行10s
	// 	time.Sleep(10 * time.Second)
	// 	close(closeChannel)
	// }()

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
