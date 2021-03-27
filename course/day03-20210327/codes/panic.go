package main

import (
	"fmt"
)

func fact(n int64) int64 {
	if n < 0 {
		panic("n不能小于0")
	}

	if n == 0 || n == 1 {
		return 1
	}
	rt := n * fact(n-1)
	return rt
}

func testError() (err error) {
	// defer 函数调用
	defer func() {
		if errMsg := recover(); errMsg != nil {
			err = fmt.Errorf("%v", errMsg)
		}
	}()

	fmt.Println(fact(-1))
	return
}

func main() {
	testError()

	fmt.Println("main")
}
