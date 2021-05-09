package main

import (
	"fmt"
	"reflect"
)

func printHi() {
	fmt.Println("hi")
}

func printHello(name string) string {
	fmt.Printf("hello : %s\n", name)
	return fmt.Sprintf("hello : %s", name)
}

func main() {
	fmt.Println("==========print HI")
	value := reflect.ValueOf(printHi)
	fmt.Println(value.Type())
	params := make([]reflect.Value, 0)
	rts := value.Call(params)
	fmt.Println(rts)

	fmt.Println("===========print Hello")
	value = reflect.ValueOf(printHello)
	fmt.Println(value.Type())
	params = make([]reflect.Value, 0)
	name := "kk"
	params = append(params, reflect.ValueOf(name))
	rts = value.Call(params)
	fmt.Println(rts[0].String())

}
