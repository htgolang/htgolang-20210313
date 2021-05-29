package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//使用map+locker实现一个线程安全的map key, value
// the map is secure when it's wrote simultaneously by other program
type MyMap struct {
	//Read and write locker protects m
	sync.RWMutex
	m map[int]int
}

func NewMyMap(n int) *MyMap {
	return &MyMap{
		m: make(map[int]int, n),
	}
}
func (m *MyMap) Get(k int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

func (m *MyMap) Set(k int, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *MyMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *MyMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}


func main() {
	num:=5
	myMap := NewMyMap(num)
	go func() {
		for i := 0; i <num; i++ {
			myMap.Delete(i)
			runtime.Gosched()
		}
	}()
	go func() {
		for i := 0; i <num; i++ {
			myMap.Set(i,i+1)
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i <num; i++ {
			myMap.Set(i,i+2)
			runtime.Gosched()
		}
	}()


	time.Sleep(3*time.Second)
	for i := 0; i <myMap.Len(); i++ {
		value,_ := myMap.Get(i)
		fmt.Printf("key:%v=>value:%v\n",i, value)
	}
}