package main

import "fmt"

func main() {
	// defer 不要写在循环之内
	// start 0 1 2 end
	// start end 0 1 2
	// start end 2 1 0
	// start end 2 2 2
	fmt.Println("start")
	// for i := 0; i < 3; i++ {
	// 	defer func() {
	// 		fmt.Println(i)
	// 	}()

	// }

	// for i := 0; i < 3; i++ {
	// 	v := i
	// 	defer func() {
	// 		fmt.Println(v)
	// 	}()

	// }

	for i := 0; i < 3; i++ {
		defer func(v int) {
			fmt.Println(v)
		}(i)

	}
	fmt.Println("end")
}
