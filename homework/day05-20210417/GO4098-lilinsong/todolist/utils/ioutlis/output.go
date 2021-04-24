package ioutlis

import (
	"fmt"
)

// Error 错误输出
func Error(text string) {
	fmt.Printf("\033[1;31;40m%s\033[0m\n", text)
}

// Success 正确输出
func Success(text string) {
	fmt.Printf("\033[1;32;40m%s\033[0m\n", text)
}

// Output 普通输出
func Output(text string) {
	fmt.Printf(text)
}
