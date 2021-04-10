package user //?

// 包名 => 在同一个文件夹内包名必须一致

import "fmt"

// 首字母小写 只能在包内使用
func Add() {
	fmt.Println("添加cmdb")
}

func Modify() {
	fmt.Println("修改cmdb")
}

func Delete() {
	fmt.Println("删除cmdb")
}

func Find() {
	fmt.Println("查询cmdb")
}
