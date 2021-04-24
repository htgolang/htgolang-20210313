package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("fprint.txt")
	fmt.Fprint(file, 1, 2, 3)
	fmt.Fprint(file, "x", "y", "z")
	fmt.Fprintln(file, 4, 5, 6)
	fmt.Fprintln(file, "a", "b", "c")
	fmt.Fprintf(file, "我的名字是: %s\n", "kk")

}
