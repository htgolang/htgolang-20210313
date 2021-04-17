package main

import (
	"fmt"
	"strconv"
	. "userManager/config"
	_ "userManager/init"
	. "userManager/utils"
)

func main() {
	/*
		在 Golang 中, 如何不使用反射实现动态类型切换, 以目前所学我无法实现,只能用
	switch / if-else 来实现. 习惯了 Python 的反射机制后不太能适应 Golang 的 "哲学",
	加个 TODO
	标记下.
	*/
	choice := Input("选择 1. TodoAdmin \n 2. UserAdmin")
	switch choice {
	case "1":
		for {
			ShowMenu(choice)
			menu := Input("输入功能序号: ")
			if menu < "0" || menu > strconv.Itoa(len(TodoFuncMap)-1) {
				fmt.Println("功能不合法, 重新输入")
				continue
			}
			TodoFuncMap[menu](TodoMap)
		}
	case "2":
		for {
			ShowMenu(choice)
			menu := Input("输入功能序号: ")
			if menu < "0" || menu > strconv.Itoa(len(UserFuncMap)-1) {
				fmt.Println("功能不合法, 重新输入")
				continue
			}
			UserFuncMap[menu](UserMap)
		}
	}
}
