package main

import (
	"fmt"
	"reflect"
)

// 处理reflect.Value类型变量v中所有的函数
func callValue(value reflect.Value) {
	fmt.Println("-------------", value.Type(), "-------------")
	//获取值对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		// 针对数组/切片类型, 对数组/切片中元素进行递归处理
		for i := 0; i < value.Len(); i++ {
			callValue(value.Index(i)) //根据索引获取数组的每个元素,并调用callValue递归调用
		}
	case reflect.Map:
		// 针对映射类型, 对映射中值进行递归处理
		iter := value.MapRange()
		for iter.Next() {
			callValue(iter.Value()) //根据切片获取数组的每个元素,并调用callValue递归调用
		}
	case reflect.Struct:
		// 针对结构体类型, 执行所有方法
		for i := 0; i < value.NumMethod(); i++ {
			callMethod(value.Method(i), value.Type(), value.Type().Method(i))
		}
	case reflect.Ptr:
		// 针对指针类型, 执行所有方法
		for i := 0; i < value.NumMethod(); i++ {
			callMethod(value.Method(i), value.Elem().Type(), value.Type().Method(i))
		}
	case reflect.Func:
		// 针对方法类型, 执行方法
		fmt.Println("Func value:", value)
		//callFunc(value)
		callFuncUseValue(value)
	default:
		fmt.Printf("Unkonw[%#v]\n", value)
	}
}

// 调用函数
func callFunc(f reflect.Value) {
	// 获取函数类型并打印
	// Type returns v's type.
	ftype := f.Type()
	fmt.Printf("value Type:%s\n", ftype.String())
	// 组装函数参数
	parameters := make([]reflect.Value, 0)
	// NumIn returns a function type's input parameter count.
	for i := 0; i < ftype.NumIn(); i++ {
		// reflect.Zero: Zero returns a Value representing the zero value for the specified type.
		// ftype.In(i): In returns the type of a function type's i'th input parameter.
		//fmt.Println("ftype.In(i):",ftype.In(i))
		parameters = append(parameters, reflect.Zero(ftype.In(i))) //通过reflect.Zero创建参数同类型零值
	}
	//调用函数并打印执行结果
	if ftype.IsVariadic() {
		// CallSlice calls the variadic function f with the input arguments parameters,
		// assigning the slice in[len(parameters)-1] to f's final variadic argument.
		fmt.Printf("%#v", f.CallSlice(parameters))
	} else {
		//调用函数
		f.Call(parameters) // 1 2 3
		//打印返回值
		//fmt.Println(f.Call(parameters))
	}
}

// 调用函数
func callFuncUseValue(f reflect.Value) {
	ftype := f.Type()
	fmt.Printf("value Type:%s\n", ftype.String())
	// parameters := make([]reflect.Value, ftype.NumIn())
	parameters := make([]reflect.Value, 0)
	fmt.Println("ftype.NumIn():", ftype.NumIn())
	for i := 0; i < ftype.NumIn(); i++ {
		parameters = append(parameters, reflect.Zero(ftype.In(i))) //通过reflect.Zero创建参数同类型零值
	}
	//调用函数并打印执行结果
	if ftype.IsVariadic() {
		fmt.Printf("%#v", f.CallSlice(parameters))
	} else {
		//调用函数
		//fmt.Printf("%#v\n", reflect.ValueOf(parameters[0]))

		// fmt.Printf("%#v\n", reflect.ValueOf(parameters[0]))
		// reflect.ValueOf(parameters[0]).SetString("127.0.0.1")
		parameters[0] = reflect.ValueOf("127.0.0.x")
		//parameters[0].SetString("127.0.0.1")
		//parameters[1].SetInt(8080)
		//
		f.Call(parameters) // 1 2 3
		//传入参数

		//打印返回值
		//fmt.Println(f.Call(parameters))
	}
}

// 调用方法
func callMethod(f reflect.Value, t reflect.Type, m reflect.Method) {
	// 获取函数类型并打印
	ftype := f.Type()
	fmt.Printf("method: %s.%s => %s\n", t.Name(), m.Name, ftype.String())

	// 组装函数参数
	parameters := make([]reflect.Value, 0)
	for i := 0; i < ftype.NumIn(); i++ {
		parameters = append(parameters, reflect.New(ftype.In(i)).Elem()) //通过reflect.New创建参数同类型零值的指针
	}
	//调用函数并打印执行结果
	if ftype.IsVariadic() {
		fmt.Println(f.CallSlice(parameters))
	} else {
		fmt.Println(f.Call(parameters))
	}
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}
type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	fmt.Printf("%s:%d\n", ip, port)
	return &Connection{Address: Address{ip, port}}
}

// 定义User结构体,并为每个属性定义标签
type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float32 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float32, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

func main() {
	//如何传参调用?
	var funcV11 func(interface{}) error = func(x interface{}) error { fmt.Println(x); return nil }
	callValue(reflect.ValueOf(funcV11))
	/*
		------------- func(interface {}) error -------------
		Func value: 0xe4b7e0
		value Type:func(interface {}) error
		123 问题1:反射函数调用,怎么传入x,比如指定x为123,这里输出为123?
	*/

	//var funcV2 func(string, int) *Connection = NewConnection
	//callValue(reflect.ValueOf(funcV2))
	/*
		------------- func(string, int) *main.Connection -------------
		Func value: 0x4cb6e0
		value Type:func(string, int) *main.Connection
		:0  问题2:反射函数调用,怎么传入参数呢?比如这里应该输出为  "192.168.1.6":8080 ,参数ip和port如何指定?
	*/
}
