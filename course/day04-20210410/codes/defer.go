package main

import (
	"fmt"
)

func e() (err error) {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			err = fmt.Errorf("%s", errMsg)
		}
	}()
	//
	// panic("test")
	return
}

func test(f bool) {
	if f {
		return
	}
	defer func() { // 注册
		fmt.Println(1)
	}()

	defer fmt.Println(2)
}

func main() {
	test(false) // defer 执行

	test(true) // defer 不执行
}
