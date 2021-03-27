package main

func test() {
	p1 := 1
	p2 := 2
	p3 := []int{1, 2, 3}
	p4 := map[string]int{"1": 1, "2": 2}

	println(&p1, &p2, &p3, &p4, p3, p4)
}

func test2() (*int, []int) {
	p1 := 1
	p2 := 2
	p3 := []int{1, 2, 3}
	p4 := map[string]int{"1": 1, "2": 2}

	println(&p1, &p2, &p3, &p4, p3, p4)
	return &p1, p3
}

func main() {
	test()
	test2()
}
