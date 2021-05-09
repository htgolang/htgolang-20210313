package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Addr json.RawMessage
}

func main() {
	txt := `{"id":100,"name":"kk","addr":{"street":"陕西省西安市","no":"10002"}}`

	var u User
	err := json.Unmarshal([]byte(txt), &u)
	fmt.Println(err, u)
}
