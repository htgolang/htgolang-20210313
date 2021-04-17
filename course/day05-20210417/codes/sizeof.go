package main

import (
	"fmt"
	"unsafe"
)

// 连续的内存
// start end
// id(s1, e1)->name(s2, e2)->tel(s3, e3)->email(s4, e4)
// 占用字节小->大
type User struct {
	id    int
	name  string
	tel   string
	email string
}

func main() {
	var i int
	fmt.Println("int:", unsafe.Sizeof(i))
	var s string
	fmt.Println("string:", unsafe.Sizeof(s))

	var u User

	fmt.Println("user:", unsafe.Sizeof(u))
}
