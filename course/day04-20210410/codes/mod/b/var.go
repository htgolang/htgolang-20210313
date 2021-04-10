package b

import "fmt"

const Name = "b"

func Call() {
	fmt.Println("b")
}

func init() {
	fmt.Println("b.init")
}

func init() {
	fmt.Println("b.init.1")
}

func init() {
	fmt.Println("b.init.2")
}
