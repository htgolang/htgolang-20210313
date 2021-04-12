package main

import (
	"fmt"
	"github.com/princebot/getpass"
	"user/commands"
	"strings"
	_ "user/init"
	"crypto/md5"
	"os"
)

const Password = "123" //md5

//Input the information
func input(prompt string) string {
	fmt.Println(prompt)
	var txt string
	fmt.Scan(&txt)
	return txt
}


func prompt() {
	prompts := commands.Prompts
	fmt.Println(strings.Repeat("*",38))
	fmt.Println("Welcome to the user management system: ")
	fmt.Println("1 exit")
	for _,v := range prompts{
		fmt.Printf("%s %s\n",v[0],v[1])
	}
	fmt.Println(strings.Repeat("*",38))
}

func main() {
	//Welcome
	//If failed three times password, then exit

	//Enter password and check the md5 value
	var count int = 0
	for {  
		passwordBytes, _ := getpass.Get("Pls enter a password: ")
		fmt.Println()
		if md5.Sum([]byte(string(passwordBytes))) != md5.Sum([]byte(Password)){
			fmt.Println("Wrong password!!!")
			fmt.Println()
			count++
		}else{
			break
		}
		if count == 3 {
			fmt.Println("You have entered a wrong password for 3 consecutive times.")
			os.Exit(1)
		}
}
	//passwordByptes --> md5 --> compare

	//for loop
	//Prompt: ==> 1: add, 2: edit, 3: del, 4: query, 5: exit

	//Store the cmd in map
	/*
	commands := map[string]func(){
		"1": user.Add,
		"2": user.Edit,
		"3": user.Del,
		"4": user.Query,
	}
	*/
	//Optimize
	commands := commands.Commands

END:
	for {
		prompt()
		cmd := input("Pls input the order: ")
		if command, ok := commands[cmd]; ok {
			command()
		}else if cmd == "1"{
			break END
		}else{
			fmt.Println("Wrong order!")
		}
	}
}