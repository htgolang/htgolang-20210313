package main

import (
	"fmt"
	"reflect"
)

type User struct{}

func (u User) TableName() string {
	return "aaa"
}

func main() {
	value := reflect.ValueOf(new(User))
	methodValue := value.MethodByName("TableName")
	if !methodValue.IsValid() {
		fmt.Println(value.Elem().Type().Name())
		return
	}
	methodType := methodValue.Type()
	if 0 != methodType.NumIn() || 1 != methodType.NumOut() {
		fmt.Println(value.Elem().Type().Name())
		return
	}

	if outType := methodType.Out(0); outType.Kind() != reflect.String {
		fmt.Println(value.Elem().Type().Name())
		return
	}

	params := make([]reflect.Value, 0)
	rts := methodValue.Call(params)
	fmt.Println(rts)
	fmt.Println(rts[0].String())
}
