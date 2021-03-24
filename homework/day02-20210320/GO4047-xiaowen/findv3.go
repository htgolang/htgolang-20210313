package main
import (
	"fmt"
)

func main() {
	cap := []int{9, 8, 19, 10, 2, 8}
	for i := 0; i < len(cap); i++ {
		for ii := 0; ii < len(cap) -1 ; ii++{
			if cap[ii] < cap[ii+1] {
				cap[ii], cap[ii+1] = cap[ii+1], cap[ii]
			}
		}
	}
	ns := []int{}
	ns = append(ns, cap[1:]...)
	ns = append(ns, cap[:1]...)
	fmt.Println(ns)
}