package main

import (
	"sync"
)

type SafeSlice struct {
	locker sync.RWMutex
	entry  []interface{}
}

func NewSafeSlice() *SafeSlice {
	return &SafeSlice{
		locker: sync.RWMutex{},
		entry:  []interface{}{},
	}
}

func (s *SafeSlice) Get(i int) interface{} {
	s.locker.RLock()
	defer s.locker.RUnlock()
	if i < 0 || i > len(s.entry) {
		panic("index out of range")
	}
	return s.entry[i]
}

func (s *SafeSlice) Set(i int, v interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()
	if i < 0 || i > len(s.entry) {
		panic("index out of range")
	}
	s.entry[i] = v
}

func (s *SafeSlice) Append(v interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.entry = append(s.entry, v)
}

func (s *SafeSlice) Delete(i int) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.entry = append(s.entry[:i], s.entry[i+1:]...)
}

func main() {
// 	slice := NewSafeSlice()
// 	wg := &sync.WaitGroup{}
// 	for i := 0; i < 10; i++ {
// 		wg.Add(4)
// 		go func() {
// 			defer wg.Done()
// 			for i := 1; i <= 1000; i++ {
// 				slice.Append(i)
// 			}
// 		}()
// 		go func()
// 	}
// }
