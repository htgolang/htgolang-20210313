package main

import (
	"fmt"
	"sync"
	"time"
)




type puc struct{
	name string
	path string
}

func downLoad(name, path string){
	fmt.Printf("%v%v download\n", path, name)
}

func main(){
	N := 10
	start := time.Now()
	wg := sync.WaitGroup{}
	m := 100	
	channel := make(chan puc, m)
	//生成下载图片
	wg.Add(1)
	go func (){
		defer wg.Done()
		for i := 1; i <= m; i++ {
			channel <- puc{fmt.Sprintf("%v.jpg", i), fmt.Sprint("/data/")}
		}
		close(channel)
	}()
// END:
// 	for{
// 		select{
// 		case v,ok := <-channel:
// 			if !ok {
// 				break END
// 			} else{
// 				downLoad(v.name, v.path)
// 			}

// 		}
// 	}

	//下载图片
	for i := 0; i< N; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for v := range channel{
				downLoad(v.name, v.path)
			}
		}()
	}

	wg.Wait()
	fmt.Print(time.Now().Sub(start))	

}