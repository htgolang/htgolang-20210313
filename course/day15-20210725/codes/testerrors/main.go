package main

import (
	"errors"
	"fmt"
)

type DbError struct {
	err error
}

func (e DbError) Error() string {
	return "dbError"
}

// func (e DbError) Unwrap() error {
// 	return e.err
// }

func main() {
	var err error = errors.New("init")

	var dbErr *DbError = &DbError{err}
	//fmt.Errorf
	// errors.New
	fmt.Println(dbErr)

	wrapError := fmt.Errorf("new error:[%w]", dbErr)
	// err => DbError => wrapError

	fmt.Println(errors.Is(wrapError, err))  // 判断错误是否为err(Unwrap)
	fmt.Println(errors.As(wrapError, &err)) // 判断错误链中是否包含err
}
