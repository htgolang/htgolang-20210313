package main

import "fmt"

func main() {
	// 1 + ... + 100
	total, incr := 0, 0

START:
	total += incr
	incr++

	if incr <= 100 {
		goto START
	}
	goto END
	fmt.Println(total)

END:
	fmt.Println("total:", total)
}
