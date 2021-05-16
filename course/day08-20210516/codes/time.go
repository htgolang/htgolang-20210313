package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Time()
	// Time string unixtime
	fmt.Println(time.Second * 10000)
	d, err := time.ParseDuration("2h46m41s")
	fmt.Println(d.Seconds(), d.Minutes(), err)

}
