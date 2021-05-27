package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	locker sync.RWMutex
	entry  map[interface{}]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		locker: sync.RWMutex{},
		entry:  make(map[interface{}]interface{}),
	}
}

func (s *SafeMap) Get(k interface{}) (interface{}, bool) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	v, ok := s.entry[k]
	return v, ok
}

func (s *SafeMap) Set(k, v interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.entry[k] = v
}

func (s *SafeMap) Delete(k interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()

	delete(s.entry, k)
}

func main() {
	wg := &sync.WaitGroup{}
	m := NewSafeMap()
	// m := make(map[int]int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for i := n * 100; i < n*1000; i++ {
				m.Set(i, i)
				// m[i] = i
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(m.entry)
	// fmt.Println(m)
}
