package main

import (
	"encoding/json"
	"fmt"
)

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
	u := User{1, "kk", Addr{"陕西省西安市", "10001"}}

	b, _ := json.MarshalIndent(u, "", "\t")
	fmt.Println(string(b))

	jsonB := `
	{
        "id": 3,
        "name": "kk",
        "addr": {
                "street": "陕西省西安市",
                "no": "10002"
        }
	}
`
	var rs User
	json.Unmarshal([]byte(jsonB), &rs)
	fmt.Println(rs)
}
