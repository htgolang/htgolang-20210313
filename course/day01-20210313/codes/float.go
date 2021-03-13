package main

import "fmt"

func main() {
	var (
		f1 float64 = 3.1
		f2 float64 = 1e-5
	)
	fmt.Printf("%T, %v\n", f1, f2)
	fmt.Printf("%f, %f\n", f1, f2)
	fmt.Printf("%e, %e\n", f1, f2)
	fmt.Printf("%g, %g\n", f1, f2)
	// %f 修饰
	// %n.mf => n占位 m小数点保留位数
	fmt.Printf("%5.3f, %5.3f\n", f1, f2)
}
