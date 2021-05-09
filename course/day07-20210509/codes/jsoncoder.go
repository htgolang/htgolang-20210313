package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
Encoder
Decoder
*/

type Addr struct {
	Street string `json:"street"`
	No     string `json:"no"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Addr Addr   `json:"addr"`
}

func main() {
	// u := User{1, "kk", Addr{"陕西省西安市", "10001"}}

	// file, _ := os.Create("user.json")
	// defer file.Close()
	// encoder := json.NewEncoder(file)
	// err := encoder.Encode(u)
	// fmt.Println(err)

	file, _ := os.Open("user.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	var user User
	err := decoder.Decode(&user)
	fmt.Println(err, user)
}
