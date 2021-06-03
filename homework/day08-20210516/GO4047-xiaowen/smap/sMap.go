package main

import (
	"fmt"
	"strconv"
	"sync"
)


type Smap struct {
	Locker sync.Mutex
	Maps map[string]string
}

func Init() *Smap{
	return &Smap{
		Locker: sync.Mutex{},
		Maps: make(map[string]string),
	}
}

func (m *Smap) Add(user, pwd string) map[string]string{
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.Maps[user] = pwd
	return m.Maps
}

func (m *Smap) Delete(user string) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	delete(m.Maps, user)
}

func (m *Smap) Set(user,pwd string) map[string]string{
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.Maps[user] = pwd
	return m.Maps
}

func (m *Smap) Get(user string){
	m.Locker.Lock()
	defer m.Locker.Unlock()
	fmt.Println(m.Maps[user])
}


func main() {
	s := Init()
	n := 10
	nums := 1000
	wg := sync.WaitGroup{}

	//添加数据
	for i := 0; i < n; i++{
		wg.Add(1)
		go func(){
			for c := 0; c < nums; c++{
				user := "user"+ strconv.Itoa(c)
				pwd := strconv.Itoa(c)
				s.Add(user, pwd)
			}
			defer wg.Done()
		}()
		
	}
	// 修改数据
	for i := 0; i < n; i++{
		wg.Add(1)
		go func(){
			for c := 0; c < nums; c++{
				user := "user"+ strconv.Itoa(c)
				pwd := strconv.Itoa(c) + "2"
				s.Set(user, pwd)
			}
			defer wg.Done()
		}()
		
	}
	//查询数据
	for i := 0; i < n; i++{
		wg.Add(1)
		go func(){
			for c := 0; c < nums; c++{
				user := "user"+ strconv.Itoa(c)
				s.Get(user)
			}
			defer wg.Done()
		}()
		
	}
	//删除数据
	for i := 0; i < n; i++{
		wg.Add(1)
		go func(){
			for c := 0; c < nums; c++{
				user := "user"+ strconv.Itoa(c)
				s.Delete(user)
			}
			defer wg.Done()
		}()
		
	}
	fmt.Println(s.Maps)
	wg.Wait()
}