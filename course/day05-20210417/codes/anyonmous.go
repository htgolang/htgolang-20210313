package main

import "fmt"

func main() {
	// 匿名结构体 => 不定义名称 => 不能再定义对应的变量 => 只需要定义该种结构的一个变量

	var user struct {
		id   int
		name string
		addr string
		tel  string
	}

	fmt.Printf("%T\n", user)
	fmt.Printf("%#v\n", user)

	user.id = 10
	fmt.Println(user.name)
	fmt.Printf("%#v\n", user)

	user = struct {
		id   int
		name string
		addr string
		tel  string
	}{1, "kk", "西安", ""}

	fmt.Printf("%#v\n", user)

	var u2 struct {
		id   int
		name string
	} = struct {
		id   int
		name string
	}{id: 1, name: "kk"}

	fmt.Printf("%#v\n", u2)

	/*
		var v T
		var v = T{}
	*/

	var u3 = struct {
		id   int
		name string
	}{1, "kk"}
	fmt.Printf("%#v\n", u3)

	u4 := struct {
		id   int
		name string
	}{1, "kk"}
	fmt.Printf("%#v\n", u4)

	u5 := &struct {
		id   int
		name string
	}{1, "kk"}
	fmt.Printf("%T\n", u5)
	fmt.Printf("%#v\n", u5)

	// 程序 => 配置
	// template => 数据(匿名结构体) => 传递给模板

	// []T => T => 结构体 => 匿名结构体
	var users = []struct {
		id   int
		name string
	}{
		{1, "kk"},
		{2, "alinx"},
	}

	fmt.Printf("%T", users)
	fmt.Printf("%#v", users)
}
