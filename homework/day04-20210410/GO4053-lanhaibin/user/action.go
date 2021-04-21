package user

import "fmt"

func AddAction() {
	var name, addr string
	fmt.Print("请输入用户名:")
	fmt.Scan(&name)
	fmt.Print("请输入地址:")
	fmt.Scan(&addr)

	Add(name, addr)
	fmt.Printf("%#v\n", users)
}

func DeleteAction() {
	var id string
	fmt.Printf("请输入要删除用户的id:")
	fmt.Scan(&id)
	Delete(id)
	fmt.Println(users)
}

func QueryAction() {
	var s string
	fmt.Print("请输入查询字符串:")
	fmt.Scan(&s)
	us, err := Query(s)
	if err != nil {
		return
	}
	fmt.Println("查询到以下用户信息:")
	for _, user := range us {
		fmt.Printf("id: %s\n", user["id"])
		fmt.Printf("用户名: %s\n", user["name"])
		fmt.Printf("地址: %s\n", user["addr"])
		fmt.Println()
	}
}

func ModifyAction() {
	var id string
	fmt.Printf("请输入与要修改用户的id:")
	fmt.Scan(&id)
	user, err := Get(id)
	if err != nil {
		fmt.Println("用户不存在!")
		return
	}
	var confirm string
	fmt.Println("请确认用户信息:")
	fmt.Printf("id: %s\n", user["id"])
	fmt.Printf("用户名: %s\n", user["name"])
	fmt.Printf("地址: %s\n", user["addr"])

	fmt.Print("是否修改?(y/n)")
	fmt.Scan(&confirm)

	if confirm == "n" || confirm == "N" {
		return
	}
	var name, addr string
	fmt.Print("请输入新的用户名:")
	fmt.Scan(&name)
	fmt.Printf("请输入新的地址:")
	fmt.Scan(&addr)
	Modify(id, name, addr)
	fmt.Println(users)
}
