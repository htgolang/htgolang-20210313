package main

import (
	"fmt"
	mt "user/commands"
	_ "user/init"
)

func main() {

	if err := mt.Login(); !err {
		fmt.Println("尝试次数过多退出程序")
		return
	}

	for {
		mt.Entrance()
		mt.MenuSelection()
	}
}
