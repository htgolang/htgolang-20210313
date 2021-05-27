package main

import (
	"fmt"
	"sync"
	"time"
)

type File struct {
	Name string
	Path string
}

func download(name, path string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("down %s to %s\n", name, path)
}

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	channel := make(chan File, 100)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(channel)
		for i := 1; i <= 100; i++ {
			file := File{
				Name: fmt.Sprintf("%d.jpg", i),
				Path: fmt.Sprintf("downloads/%d.jpg", i),
			}
			channel <- file
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range channel {
				download(file.Name, file.Path)
			}
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
