package main

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

//Add user function
func addUser() error {
	name := input("Please enter a name: ")
	addr := input("Please enter an addr: ")
	tel := input("Please enter a tel number: ")

	//Check the user exist or not
	for _, user := range users {
		u, _ := user["name"]
		if u == name {
			fmt.Println(name)
			return fmt.Errorf("The user %s exists.", name)
		}
	}

	//Add users
	users = append(users, map[string]string{
		"id":   strconv.Itoa(getUserId()),
		"name": name,
		"addr": addr,
		"tel":  tel,
	})
	return nil
}

func main() {
	if err := addUser(); err == nil {
		for _, v := range users {
			fmt.Println(v)
		}
	} else {
		fmt.Println("Add user failed.", err)
	}
}
