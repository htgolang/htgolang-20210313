package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	//The max number of times for the password attempts
	mathAuth = 3
	password = "123"
)

//public input function
func inputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input) //remove the space
}

//verify the password
func auth() bool {
	//if the password passed, return true
	var input string
	for i := 0; i < mathAuth; i++ {
		fmt.Print("Pls enter a password:")
		fmt.Scan(&input)
		if password == input {
			//success
			return true
		} else {
			//fail
			fmt.Println("Wrong password!")
		}
	}
	return false //failed three times, will be false
}

//format the output
func printUser(pu int, user map[string]string) {
	fmt.Println("ID: ", pu)
	fmt.Println("Name: ", user["name"])
	fmt.Println("Birth: ", user["birth"])
	fmt.Println("Tel: ", user["tel"])
	fmt.Println("Addr: ", user["addr"])
	fmt.Println("Desc: ", user["desc"])
}

//query user function
func queryUser(users map[int]map[string]string) {
	q := inputString("Pls enter a query string: ")
	fmt.Println("==========================")
	for k, v := range users {
		//if contain the related string, then print
		if strings.Contains(v["name"], q) || strings.Contains(v["tel"], q) || strings.Contains(v["add"], q) || strings.Contains(v["desc"], q) {
			printUser(k, v)
			fmt.Println("---------------------------")
		}
	}
	fmt.Println("==========================")
}

//public input user info
func inputUser() map[string]string {
	user := map[string]string{}
	user["name"] = inputString("Pls enter a name: ")
	user["birthday"] = inputString("Pls enter the birthday: ")
	user["tel"] = inputString("Pls enter the phone number: ")
	user["addr"] = inputString("Pls enter an address: ")
	user["desc"] = inputString("Pls enter the desc: ")
	return user //return map
}

//get the user id
func getId(users map[int]map[string]string) int {
	var id int
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

//add user function
func addUser(users map[int]map[string]string) {
	id := getId(users)
	//add
	user := inputUser()
	users[id] = user
	fmt.Println("add success.")
}

//modify user function
func modifyUser(users map[int]map[string]string) {
	idString := inputString("Pls enter the modify user id: ")
	//string to int
	if id, err := strconv.Atoi(idString); err == nil {
		//key in map or not
		if user, ok := users[id]; ok {
			fmt.Println("The modify info: ")
			printUser(id, user)
			input := inputString("confirm(y/n)? ")
			if input == "y" || input == "Y" {
				user := inputUser()
				users[id] = user
				fmt.Println("modify success.")
			}
		} else {
			fmt.Println("User id not exist.")
		}
	} else {
		fmt.Println("Wrong id.")
	}
}

//delete user function
func deleteUser(users map[int]map[string]string) {
	//similar with modify
	idString := inputString("Pls enter the delete user id: ")
	//string to int
	if id, err := strconv.Atoi(idString); err == nil {
		//key in map or not
		if user, ok := users[id]; ok {
			fmt.Println("The delete info: ")
			printUser(id, user)
			input := inputString("confirm(y/n)? ")
			if input == "y" || input == "Y" {
				//delete key
				delete(users, id)
				fmt.Println("remove success.")
			}
		} else {
			fmt.Println("User id not exist.")
		}
	} else {
		fmt.Println("Wrong id.")
	}
}

func main() {
	if !auth() {
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
	//detail: id, name, birth, tel, addr, desc
	users := map[int]map[string]string{}

	callbacks := map[string]func(map[int]map[string]string){
		"1": queryUser,
		"2": addUser,
		"3": modifyUser,
		"4": deleteUser,
		"5": func(users map[int]map[string]string) {
			os.Exit(0)
		},
	}

	fmt.Println("Welcome to user management system!")
	for {
		fmt.Println(menu_list)
		if callback, ok := callbacks[inputString("Pls enter the command: ")]; ok {
			callback(users)
		} else {
			fmt.Println("Wrong command.")
		}
	}
}
