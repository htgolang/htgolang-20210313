package main

import "fmt"

func sayHello() {

}

func add23() int {
	return 2 + 3
}

func calc() (int, int, int, int) {
	return 2 + 3, 2 - 3, 2 * 3, 2 / 3
}

func returnType() (int, bool, string, []int) {
	return 1, true, "xxx", nil
}

// 命名返回值

// 将函数执行完成时name的返回
func returnName() (name string) {

	name = "kkxcxxx"
	name = "xxxxx"
	name = "aaa"
	name += "111"
	return
}

func returnNameArgs() (name string, isBody bool) {
	name = name + "kk"
	return
}

func returnMergeArgs() (b bool, n1, n2 string, isBody bool) {
	b = true
	return
}

func main() {
	sayHello()

	x := add23()
	fmt.Println(x)
	fmt.Println(add23())
	fmt.Println(calc())
	r1, r2, r3, r4 := calc()
	fmt.Println(r1, r2, r3, r4)

	returnType()
	rt1, rt2, rt3, rt4 := returnType()
	fmt.Printf("%T %T %T %T\n", rt1, rt2, rt3, rt4)
	fmt.Println(rt1, rt2, rt3, rt4)

	fmt.Printf("%q\n", returnName())
	fmt.Println(returnNameArgs())
	ra1, ra2 := returnNameArgs()
	fmt.Printf("%T %T %v %v\n", ra1, ra2, ra1, ra2)

	rm1, rm2, rm3, rm4 := returnMergeArgs()
	fmt.Printf("%T %T %T %T\n", rm1, rm2, rm3, rm4)
	fmt.Printf("%v %q %q %v\n", rm1, rm2, rm3, rm4)
}
