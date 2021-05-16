package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享数据
var money int = 500 * 10000

// 公司可以让大家借钱(提前支出) 还钱
// 账户 => 500 * 10000
// 每个人都可以借钱 => 还钱

// 解决方法： 借钱和还钱过程中不要终端
// 加锁

var locker sync.Mutex

// 借钱
func borrow(m int) int {
	locker.Lock()
	defer locker.Unlock()

	if money < m {
		// 释放锁
		// locker.Unlock()
		return 0
	}
	money -= m

	// locker.Unlock()

	return m
}

// 还钱
func repay(m int) {
	locker.Lock()
	money += m
	locker.Unlock()
}

func main() {
	// 2 个人借钱
	// A, B

	wg := sync.WaitGroup{}

	for person := 'A'; person <= 'Z'; person++ {
		wg.Add(1)
		go func(person rune) {
			// func(person rune) {
			// 借1000次 借2000
			m := borrow(2000)
			time.Sleep(2 * time.Microsecond)
			// 还1000次
			repay(m)

			wg.Done()
		}(person)
	}

	wg.Wait()
	fmt.Println("money:", money) // 500 * 10000
}

// 解决问题
// write(context) => file
// 两个人同时打开文件 => 写数据 => 文件会乱 => 加锁 os
// 数据库 =>
