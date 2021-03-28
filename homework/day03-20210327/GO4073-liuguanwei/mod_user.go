package main

import "fmt"

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
	{
		"id":   "3",
		"name": "Jerry",
		"addr": "gz",
		"tel":  "13325236831",
	},
	{
		"id":   "4",
		"name": "Linda",
		"addr": "sh",
		"tel":  "15325236888",
	},
}

//Enter the user information
func input(prompt string) string {
	var text string
	fmt.Println(prompt)
	fmt.Scan(&text)
	return text
}

func printUsers() {
	fmt.Println("The current users are: ")
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println()
}

func modUser() {
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

func main() {
	modUser()
}
