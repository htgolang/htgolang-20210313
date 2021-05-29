package main

import (
	"fmt"
	"sync"
	"time"
)

/*使用slice+locker实现一个线程安全的切片
启动4种操作 每种操作10个例程 循环1000次对应操作
添加元素
删除元素
查看元素
修改元素
mutex/rwmutex */
type MySlice struct {
	sync.RWMutex
	s []int
}

func NewMySlice(n int) *MySlice {
	return &MySlice{
		s:make([]int,n),
	}
}
func (m *MySlice) Add(v int) {
	m.Lock()
	defer m.Unlock()
	m.s = append(m.s,v)
}

func (m *MySlice) Delete(index int) {
	m.Lock()
	defer m.Unlock()
	if len(m.s)-1 >0{
		copy(m.s[index:], m.s[index+1:])
		m.s = m.s[:len(m.s)-1]
	}else if len(m.s)-1==0{
		m.s = []int{}
	}
}

func (m *MySlice) modify(index ,value int) {
	m.Lock()
	defer m.Unlock()
	m.s[index] = value
}

func (m *MySlice) show() []int{
	m.Lock()
	defer m.Unlock()
	return m.s
}

func (m *MySlice) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.s)
}


func main() {
	mySlice := NewMySlice(0)
	num:=2
	goNum := 8
	for n := 0; n < goNum; n++ {
		go func() {
			for i := 0; i < num; i++ {
				mySlice.Add(i+5)
			}
		}()
	}
	time.Sleep(5*time.Second)
	fmt.Println(mySlice.show())

	for m := 0; m < goNum; m++ {
		go func() {
			for i := 0; i < num; i++ {
				mySlice.Delete(i)
			}
		}()
	}
	time.Sleep(5*time.Second)
	fmt.Println(mySlice.show())

	for q := 0; q < len(mySlice.s); q++ {
		mySlice.modify(q,q+1)
	}
	fmt.Println(mySlice.show())
}
/*
[5 6 5 6 5 6 5 6 5 6 5 6 5 6 5 6]
[]
[]
*/