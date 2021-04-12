package user

import (
	"fmt"
	"strconv"
)

//User Information
var users = []map[string]string{
	{
		"id":   "1",
		"name": "Mike",
		"addr": "bj",
		"tel":  "15325236830",
	},
	{
		"id":   "2",
		"name": "Tom",
		"addr": "sz",
		"tel":  "19525236831",
	},
}

//Get the user id
func getUserId() int {
	id := 0
	for _, userid := range users {
		i, _ := strconv.Atoi(userid["id"])
		if i > id {
			id = i
		}
	}
	return id + 1
}

//Enter the user information
func input(prompt string) string {
	var text string
	fmt.Println(prompt)
	fmt.Scan(&text)
	return text
}

//Print users
func printUsers() {
	fmt.Println("The current users are: ")
	for _, user := range users {
		if len(user) == 0{
			continue
		}
		fmt.Println(user)
	}
	fmt.Println()
}

func Add() {
	name := input("Please enter a name: ")
	addr := input("Please enter an addr: ")
	tel := input("Please enter a tel number: ")

	//Check the user exist or not
	for _, user := range users {
		u, _ := user["name"]
		if u == name {
			fmt.Println(name)
			fmt.Printf("The user %s exists.\n", name)
		}
	}

	//Add users
	users = append(users, map[string]string{
		"id":   strconv.Itoa(getUserId()),
		"name": name,
		"addr": addr,
		"tel":  tel,
	})
}

func Edit(){
	id := input("Please enter the modify id: ")
	for _, user := range users {
		if user["id"] == id {
			fmt.Println(user)
			var confirm string
			var modValue string
			var modV string
			fmt.Println("Are you confirm to modify?[y/yes/Y/YES]")
			fmt.Scan(&confirm)
			if confirm == "y" || confirm == "yes" || confirm == "Y" || confirm == "YES" {
				fmt.Printf("Please enter choose the value: ")
				fmt.Scan(&modValue)
				if modValue == "name" {
					fmt.Println("Pls enter a new name: ")
					fmt.Scan(&modV)
					user[modValue] = modV
				} else if modValue == "tel" {
					fmt.Println("Pls enter a new tel number: ")
					fmt.Scan(&modV)
					user[modValue] = modV
				} else if modValue == "addr" {
					fmt.Println("Pls enter a new address: ")
					fmt.Scan(&modV)
					user[modValue] = modV
				} else {
					fmt.Println("==>Wrong value.")
					return
				}
				printUsers()
				return
			}
		}
	}
	fmt.Println("The user is not exist.")
}

func Del() {
	id := input("Please enter the del id: ")
	for _, user := range users {
		if user["id"] == id {
			fmt.Println(user)
			var confirm string
			fmt.Println("Are you confirm?[y/yes/Y/YES]")
			fmt.Scan(&confirm)
			if confirm == "y" || confirm == "yes" || confirm == "Y" || confirm == "YES" {
				fmt.Printf("The user %s is deleted.\n\n", user["name"])
				delete(user, "id")
				delete(user, "name")
				delete(user, "addr")
				delete(user, "tel")
				printUsers()
				return
			}
		}
	}
	fmt.Println("The user is not exist.")
}

func Query() {
	var str string
	fmt.Println("Please enter a query string: ")
	fmt.Scan(&str)
	for _, user := range users {
		if user["id"] == str || user["name"] == str || user["addr"] == str || user["tel"] == str {
			fmt.Println(user)
			fmt.Println()
			break
		}
	}
}
