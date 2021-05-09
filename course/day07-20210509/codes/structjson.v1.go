package main

import (
	"encoding/json"
	"fmt"
)

// 需要设置属性与Json字符串中key的对应关系
type User struct {
	Id     int
	Name   string
	IsBoy  bool
	Scores []float64
	Phone  map[string]string
}

func main() {
	user := User{1, "kk", true, []float64{1, 3, 4}, map[string]string{"tel": "1520000000"}}

	b, err := json.Marshal(user)
	fmt.Printf("%#v, %#v\n", err, string(b))

	jsonB := `
	{
		"Id": 10,
		"Name" : "kk",
		"IsBoy" : false,
		"Phone" : {"mobile" : "123456"}
	}
	`
	var u User
	err = json.Unmarshal([]byte(jsonB), &u)
	fmt.Printf("%#v, %#v\n", err, u)

}
