package main

import (
	"fmt"
)

// 管道
func main() {
	// CSP原理(顺序通信进程)
	// 队列 buffer区（线程/例程/协程安全的） -> 入队列 -> 出队列
	// 类型 chan T T Go无限制  常常放:Go 值类型
	// 如果放一个切片到管道，其他例程从管道中读取切片后对切片元素修改，会对放入端切片有影响

	// 放置int类型元素的管道
	var channel chan int

	fmt.Printf("%T, %#v\n", channel, channel)

	// 赋值 make函数
	// 对应 有对管道不同的读和写操作的例程
	// make(chan T) 不带缓冲区的管道
	// 写入 有对应的读
	// 读 有对对应的写

	// make(chan T, size) 带缓冲区的管道, 长度size
	// 5 size => 1 1 1 1 1 => 1 等待 如果go检查无goroutine可以读, 死锁
	// 读 读空channel, 如果go检查入goroutine可以写, 死锁

	// kk要一个苹果 没有人给kk找苹果（goroutine）
	// kk给大家吃苹果 放在篮子中(5) 1 2 3 4 5 6没有人吃苹果
	channel = make(chan int)

	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	channel <- 4
	// }()

	// // 如何对channel写@@
	// // 如何对channel读
	// e := <-channel // channel()
	// fmt.Println(e)

	go func() {
		fmt.Println(<-channel)
	}()

	channel <- 1

}
