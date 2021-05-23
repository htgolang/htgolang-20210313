package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	var locker sync.Mutex
	stats := make(map[string]string)
	wg.Add(30)
	for i := 0; i < 10; i++ {
		//add
		go func() {
			for num := 0; num < 1000; num++ {
				locker.Lock()
				newNum := "reid" + strconv.Itoa(num)
				stats[newNum] = strconv.Itoa(num)
				locker.Unlock()
			}
			wg.Done()
		}()

		//mod
		go func() {
			for num := 0; num < 1000; num++ {
				locker.Lock()
				newNum := "reid" + strconv.Itoa(num)
				newValue := "rr" + strconv.Itoa(num)
				stats[newNum] = newValue
				locker.Unlock()
			}
			wg.Done()
		}()

		//del
		go func() {
			for num := 0; num < 1000; num++ {
				locker.Lock()
				newNum := "reid" + strconv.Itoa(num)
				delete(stats, newNum)
				locker.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("The final result: ", stats)
}
