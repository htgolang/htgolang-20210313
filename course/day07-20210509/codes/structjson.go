package main

import (
	"encoding/json"
	"fmt"
)

// 需要设置属性与Json字符串中key的对应关系
// Id <=> json.pk
// 结构体的(属性)标签 ``字符串 key1:"value1" key2:"value2"
// json:"name"
// json:"-", 序列化和反序列化时忽略属性
// json:"name,omitempty"

// json:"name,string,omitempty"
// json:",string,omitempty"
// json:"name,string"
// omitempty 属性对应的值未零值在json字符串中不包含
type User struct {
	Id       int               `json:"pk,string"`
	Name     string            `json:"name"`
	Password string            `json:"-"`
	IsBoy    bool              `json:"isBoy,string,omitempty"`
	Scores   []float64         `json:"scores"`
	Phone    map[string]string `json:"phone"`
}

func main() {
	user := User{1, "kk", "123@456", true, []float64{1, 3, 4}, map[string]string{"tel": "1520000000"}}

	b, err := json.Marshal(user)
	fmt.Printf("%#v, %#v\n", err, string(b))

	jsonB := `
	{
		"pk": "10",
		"name" : "kk",
		"Password":"123@abc",
		"isBoy" : "false",
		"phone" : {"mobile" : "123456"}
	}
	`
	var u User
	err = json.Unmarshal([]byte(jsonB), &u)
	fmt.Printf("%#v, %#v\n", err, u)

	users := []User{
		{1, "kk", "123@4561", true, []float64{1, 3, 4}, map[string]string{"tel": "1520000000"}},
		{2, "kk2", "123@4562", true, []float64{1, 3, 4}, map[string]string{"tel": "1520000000"}},
		{3, "kk3", "123@4563", true, []float64{1, 3, 4}, map[string]string{"tel": "1520000000"}},
	}

	b, err = json.MarshalIndent(users, "", "\t")
	fmt.Printf("%#v, %s\n", err, string(b))

	jsonB = `
	[
        {
                "pk": "11",
                "name": "kk",
                "isBoy": "true",
                "scores": [
                        1,
                        3,
                        4
                ],
                "phone": {
                        "tel": "1520000000"
                }
        },
        {
                "pk": "12",
                "name": "kk2",
                "isBoy": "true",
                "scores": [
                        1,
                        3,
                        4
                ],
                "phone": {
                        "tel": "1520000000"
                }
        },
        {
                "pk": "13",
                "name": "kk3",
                "isBoy": "true",
                "scores": [
                        1,
                        3,
                        4
                ],
                "phone": {
                        "tel": "1520000000"
                }
        }
]
	`

	var us []User
	err = json.Unmarshal([]byte(jsonB), &us)
	fmt.Println(err, us)
}
