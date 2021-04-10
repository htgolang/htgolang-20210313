package main

import (
	"fmt"
	"log"
)

func test() {
	defer func() {
		recover()
	}()
	log.Panicln("panic")
}

func testFatal() {
	log.Fatalln("fatal")
}

func main() {
	log.Print("00.我叫kk")

	fmt.Println(log.Flags())

	log.SetPrefix("前缀:") // 设置日志前缀
	log.SetFlags(log.Ldate)

	log.Println("01.我叫kk")
	log.SetFlags(log.Ltime)

	log.Printf("日志: %s", "02.我叫kk")

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	test()
	testFatal()
	log.Println("over")
}
