package main

import "fmt"

func main() {
	var isBody bool
	fmt.Printf("%T, %v\n", isBody, isBody)

	t1, t2 := true, true
	f1, f2 := false, false
	fmt.Println(t1 && t2, t1 && f1, f1 && t2, f1 && f2)
	// true, false, false, false
	fmt.Println(t1 || t2, t1 || f1, f1 || t2, f1 || f2)
	// true, true, true, false
	fmt.Println(!t1, !f1)
	// false, true

	// if, for
	fmt.Println(f1 == f2, f1 == t1)

	fmt.Printf("%t, %t\n", t1, f1)
}
