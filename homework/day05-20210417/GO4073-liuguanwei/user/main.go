package main

import (
	"fmt"
	"strings"
	"user/commands"
	_ "user/init"
	"user/user"
)

//Input the information
func input(prompt string) string {
	fmt.Println(prompt)
	var txt string
	fmt.Scan(&txt)
	return txt
}

func prompt() {
	prompts := commands.Prompts
	fmt.Println(strings.Repeat("*", 38))
	fmt.Println("Welcome to the user management system: ")
	fmt.Println("1 exit")
	for _, v := range prompts {
		fmt.Printf("%s %s\n", v[0], v[1])
	}
	fmt.Println(strings.Repeat("*", 38))
}

func main() {
	//Welcome
	//If failed three times password, then exit

	//Enter password and check the md5 value
	if !user.Auth() {
		fmt.Println("Password failed three times, exit.")
		return //will exit
	}

	commands := commands.Commands

END:
	for {
		prompt()
		cmd := input("Pls input the order: ")
		if command, ok := commands[cmd]; ok {
			command()
		} else if cmd == "1" {
			break END
		} else {
			fmt.Println("Wrong order!")
		}
	}
}
