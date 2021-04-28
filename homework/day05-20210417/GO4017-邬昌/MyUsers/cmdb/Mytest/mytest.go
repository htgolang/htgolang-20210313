package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
}

func main() {

	var stu Student
	stu.name = "张三"

	test(stu)

}

func test(s Student) {

	fmt.Println(s.name)
}
