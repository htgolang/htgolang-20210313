package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var Id int64 = 0

var users []map[string]string

func Get(id string) (map[string]string, error) {
	for _, user := range users {
		if user["id"] == id {
			return user, nil
		}
	}
	return map[string]string{}, errors.New("")
}

func Add(name, addr string) {
	Id += 1
	id := strconv.FormatInt(Id, 10)
	user := map[string]string{
		"id":   id,
		"name": name,
		"addr": addr,
	}

	users = append(users, user)
}

func Delete(id string) {
	for i, user := range users {
		if user["id"] == id {
			fmt.Println("请确认用户信息:")
			fmt.Printf("id: %s\n", user["id"])
			fmt.Printf("name: %s\n", user["name"])
			fmt.Printf("addr: %s\n", user["addr"])
			var confirm string
			fmt.Printf("是否删除?(y/n)")
			fmt.Scan(&confirm)
			if confirm == "y" || confirm == "Y" {
				users = append(users[:i], users[i+1:]...)
				fmt.Println("删除成功!")
			}
		}
	}
}

func Query(s string) ([]map[string]string, error) {
	var ret []map[string]string
	for _, user := range users {
		if strings.Contains(user["id"], s) || strings.Contains(user["name"], s) || strings.Contains(user["addr"], s) {
			ret = append(ret, user)
		}
	}
	if len(ret) == 0 {
		err := errors.New("未查询到任何用户信息!")
		return ret, err
	} else {
		return ret, nil
	}
}

func Modify(id, newname, newaddr string) {
	for _, user := range users {
		if id == user["id"] {
			user["name"] = newname
			user["addr"] = newaddr
			break
		}
	}
}
