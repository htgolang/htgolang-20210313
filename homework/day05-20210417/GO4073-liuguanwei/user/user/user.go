package user

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
)

const (
	//The max number of times for the password attempts
	MathAuth = 3
	password = "202cb962ac59075b964b07152d234b70" //md5 32bit result
)

//define the struct
type User struct {
	ID   int
	Name string
	Tel  string
	Addr string
}

//redefine user
var users map[int]User

func init() {
	users = make(map[int]User)
}

//optimize the print function
func (u User) String() string {
	return fmt.Sprintf("ID: %d\nName:%s\nAddr:%s\nPhone:%s\n", u.ID, u.Name, u.Addr, u.Tel)
}

//verify the password
func Auth() bool {
	//if the password passed, return true
	// var input string
	for i := 0; i < MathAuth; i++ {
		fmt.Print("Pls enter a password: ")
		// fmt.Scan(&input)
		bytes, _ := gopass.GetPasswd()
		// from string to []byte
		pass := fmt.Sprintf("%x", md5.Sum(bytes))
		if password == pass {
			//success
			return true
		} else {
			//fail
			fmt.Println("Wrong password!")
		}
	}
	return false //failed three times, will be false
}

//Get the user id
func getUserId() int {
	var id int
	for userid := range users {
		if id < userid {
			id = userid
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
	for id, user := range users {
		if id == 0 {
			continue
		}
		fmt.Println(user)
	}
	fmt.Println()
}

//public input user info
func inputUser(id int) User {
	var user User
	user.ID = id
	user.Name = input("Pls enter a name: ")
	user.Tel = input("Pls enter the phone number: ")
	user.Addr = input("Pls enter an address: ")
	return user //return struct
}

func Add() {
	id := getUserId()
	//add
	// user := inputUser()
	user := inputUser(id)
	users[id] = user
	fmt.Println("add success.")
}

func Edit() {
	idString := input("Pls enter the modify user id: ")
	//string to int
	if id, err := strconv.Atoi(idString); err == nil {
		//key in map or not
		if user, ok := users[id]; ok {
			fmt.Println("The modify info: ")
			// printUser(id, user)
			fmt.Println(user)
			inputString := input("confirm(y/n)? ")
			if inputString == "y" || inputString == "Y" {
				// user := inputUser()
				user := inputUser(id)
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

func Del() {
	idString := input("Pls enter the delete user id: ")
	//string to int
	if id, err := strconv.Atoi(idString); err == nil {
		//key in map or not
		if user, ok := users[id]; ok {
			fmt.Println("The delete info: ")
			// printUser(id, user)
			fmt.Println(user)
			inputString := input("confirm(y/n)? ")
			if inputString == "y" || inputString == "Y" {
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

func Query() {
	q := input("Pls enter a query string: ")
	fmt.Println("==========================")
	// for k, v := range users {
	for _, v := range users {
		//if contain the related string, then print
		// if strings.Contains(v["name"], q) || strings.Contains(v["tel"], q) || strings.Contains(v["add"], q) || strings.Contains(v["desc"], q) {
		if strings.Contains(v.Name, q) || strings.Contains(v.Tel, q) || strings.Contains(v.Addr, q) {
			// printUser(k, v)
			fmt.Println(v)
			fmt.Println("---------------------------")
		}
	}
	fmt.Println("==========================")
}
