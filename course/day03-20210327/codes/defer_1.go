package main

import "fmt"

func test() (i int) {
	defer func() {
		fmt.Println("defer")
		i = 2 // 除了error场景 不要在defer中使用修改返回值
	}()

	i = 1

	return
}

func main() {
	fmt.Println(test())

}
