package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num:=5
	testMap := make(map[int]int,num)
	go func() {
		for i := 0; i <num; i++ {
			testMap[i] = i+1
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i <num; i++ {
			testMap[i] = i+2
			runtime.Gosched()
		}
	}()
	time.Sleep(5*time.Second)
	for i := 0; i <len(testMap); i++ {
		fmt.Printf("key:%v=>value:%v\n",i, testMap[i])
	}
}
/*
报错:
fatal error: concurrent map writes

goroutine 19 [running]:
runtime.throw(0x10cc013, 0x15)
        /usr/local/go/src/runtime/panic.go:1117 +0x72 fp=0xc000039758 sp=0xc000039728 pc=0x10327f2
runtime.mapassign_fast64(0x10b38e0, 0xc000100030, 0x0, 0x0)
*/