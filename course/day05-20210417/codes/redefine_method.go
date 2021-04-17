package main

import "fmt"

type Counter int

func (c *Counter) Incr() {
	(*c)++
}

func (c *Counter) Decr() {
	(*c)--
}

/*
type Counter2 = int

func (c *Counter2) Incr() {
	(*c)++
}
*/

func main() {
	var counter Counter

	counter.Incr()
	fmt.Println(counter)
	counter.Decr()
	fmt.Println(counter)
}
