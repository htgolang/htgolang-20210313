package main

import (
	"fmt"
	"os"
	userpkg "usersystem/users"
)

func main() {
	if !userpkg.Auth() {
		fmt.Println("Password failed three times, exit.")
		return //will exit
	}
	//menu list
	menu_list := `***********************************
1. query
2. add
3. mod
4. del
5. exit
***********************************`
	callbacks := map[string]func(){
		"1": userpkg.QueryUser,
		"2": userpkg.AddUser,
		"3": userpkg.ModifyUser,
		"4": userpkg.DeleteUser,
		"5": func() {
			os.Exit(0)
		},
	}

	fmt.Println("Welcome to user management system!")
	for {
		fmt.Println(menu_list)
		if callback, ok := callbacks[userpkg.InputString("Pls enter the command: ")]; ok {
			callback()
		} else {
			fmt.Println("Wrong command.")
		}
	}
}
