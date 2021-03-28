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

func queryUser() {
	for {
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
}

func main() {
	queryUser()
}
