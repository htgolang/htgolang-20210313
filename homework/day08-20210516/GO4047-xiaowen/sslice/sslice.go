package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)


type Sslice struct {
	Locker sync.Mutex
	Slices []string
}

func Init() *Sslice{
	return &Sslice{
		Locker: sync.Mutex{},
		Slices: []string{},
	}
}

func (s *Sslice) Add(user string) []string{
	s.Locker.Lock()
	defer s.Locker.Unlock()
	return append(s.Slices, user)
}

func (s *Sslice) Delete(id int) error {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	if len(s.Slices) < id || len(s.Slices) == 0{
		log.Fatalln("index out of range")
		return fmt.Errorf("%v", "index out of range")
	} else{
		fmt.Println(id, len(s.Slices))
		s.Slices = append(s.Slices[:id], s.Slices[id+1:]...)
		return nil
	}
}
	

func (s *Sslice) Set(user string, id int) ([]string,error){
	s.Locker.Lock()
	defer s.Locker.Unlock()
	if len(s.Slices) < id || len(s.Slices) == 0{
		log.Fatalln("index out of range")
		return nil,fmt.Errorf("%v", "index out of range")
	} else{
		s.Slices[id] = user
		return s.Slices,nil
	}

}

func (s *Sslice) Get(id int) error{
	s.Locker.Lock()
	defer s.Locker.Unlock()
	if len(s.Slices) < id || len(s.Slices) == 0{
		log.Fatalln("index out of range")
		return fmt.Errorf("%v", "index out of range")
	} else{
		fmt.Println(s.Slices[id])
		return nil
	}
	
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
				s.Add(user)
			}
			defer wg.Done()
		}()
		
	}
	// 修改数据

	for i := 0; i < n; i++{
		wg.Add(1)
		go func(){
			END:
			for c := 0; c < nums; c++{
				user := "user"+ strconv.Itoa(c)
				_, err := s.Set(user, c)
				if err != nil{
					break END
				}
			}
			
			defer wg.Done()
		}()

		
	}
	//查询数据
	for i := 0; i < n; i++{
		wg.Add(1)
		go func(){
			END:
			for c := 0; c < nums; c++{
				err := s.Get(c)
				if err != nil{
					break END
				}
			}
			defer wg.Done()
		}()
		
	}
	//删除数据
	for i := 0; i < n; i++{
		wg.Add(1)
		go func(){
			END:
			for c := 0; c < nums; c++{
				err := s.Delete(c)
				if err != nil{
					break END
				}
			}
			defer wg.Done()
		}()
		
	}
	wg.Wait()
}