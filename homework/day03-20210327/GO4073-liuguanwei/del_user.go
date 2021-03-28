package main

import (
	"fmt"
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
		if len(user) == 0 {
			continue
		}
		fmt.Println(user)
	}
	fmt.Println()
}

func delUser() {
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

func main() {
	delUser()
}
