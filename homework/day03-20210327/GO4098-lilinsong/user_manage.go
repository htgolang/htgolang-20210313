package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var data = []map[string]string{
	{
		"id":   "1",
		"name": "蜡笔小新",
		"addr": "春日部",
		"tel":  "123456",
	},
}

func addUser() error {
	var (
		id, name, addr, tel, txt string
		count                      int
	)
	for {
		fmt.Print("请输入用户的名称：")
		fmt.Scan(&name)
		fmt.Print("请输入地址：")
		fmt.Scan(&addr)
		fmt.Print("请输入电话：")
		fmt.Scan(&tel)
		fmt.Printf("name：%s, addr：%s, tel：%s\n", name, addr, tel)
		fmt.Print("是否确认(y/n)")
		fmt.Scan(&flag1)
		txt = strings.ToLower(strings.TrimSpace(flag1))
		if txt == "y" {
			fmt.Println("正在提交数据")
			index := len(data) + 1
			id = strconv.Itoa(index)
			data = append(data, map[string]string{
				"id":   id,
				"name": name,
				"addr": addr,
				"tel":  tel,
			})
			fmt.Println("写入数据完成")
			fmt.Printf("用户信息：\nname: %s, addr: %s, tel: %s\n",
				data[index-1]["name"], data[index-1]["addr"], data[index-1]["tel"])
		} else if txt == "n" {
			fmt.Println("请重新输入")
			continue
		} else {
			fmt.Println("输入有误从新输入")
			count++
		}
		if count > 3 {
			fmt.Println("输入错误次数过多，程序即将退出")
			return errors.New("addUser error")
		}

		fmt.Print("是否继续添加用户(y/n)")
		fmt.Scan(&txt)
		txt = strings.ToLower(strings.TrimSpace(txt))
		if txt == "y" {
			continue
		} else if txt == "n" {
			return nil
		} else {
			fmt.Println("输入有误从新输入(y/n)")
		}
	}
}

func selectData() {
	var (
		num, count int
	)
	for {
		flag := false
		fmt.Printf(`
		请输入序号
	1. 查询所有用户信息
	2. 查询用户的地址
	3. 查询用户的电话
	4. 退出
`)
		fmt.Print("请输输入序号")
		if _, err := fmt.Scan(&num); err != nil {
			fmt.Println("请输入正确的范围数字, 1~3")
			count++
			continue
		} else if num == 4 {
			break
		}

		switch num {
		case 1:
			for k := range data {
				fmt.Printf("id: %s, name: %s, addr: %s, tel: %s\n",
					data[k]["id"], data[k]["name"], data[k]["addr"], data[k]["tel"])
			}
		case 2:
			var name string
			fmt.Print("请输入查询用户的名称:")
			fmt.Scan(&name)
			strings.TrimSpace(name)
			fmt.Println("正在验证用户信息")
			for k := range data {
				if data[k]["name"] == name {
					fmt.Printf("name: %s, addr: %s\n",
						data[k]["name"], data[k]["addr"])
					flag = true
				}
			}
			if !flag {
				fmt.Println("用户不存在")
			}
		case 3:
			var name string
			strings.TrimSpace(name)
			fmt.Print("请输入查询用户的名称")
			fmt.Scan(&name)

			fmt.Println("正在验证用户信息")
			for k := range data {
				if data[k]["name"] == name {
					fmt.Printf("name: %s, addr: %s\n",
						data[k]["name"], data[k]["tel"])
					flag = true
				}
			}
			if !flag {
				fmt.Println("用户不存在")
			}
		default:
			fmt.Println("请输入正确的范围数字, 1~3")
			count++
		}
	}
}

func updateData() {
	var (
		name, content, field, txt string
	)
	fmt.Print("请输入用户的名称")
	fmt.Scan(&name)
	strings.TrimSpace(name)
	for {
		flag := false
		fmt.Println("正在验证用户信息")
		for k := range data {
			if data[k]["name"] == name {
				flag = true
				fmt.Print("请输入要修改的字段(name|addr|tel):")
				fmt.Scan(&field)
				strings.TrimSpace(field)
				if field == "name" || field == "addr" || field == "tel" {
					fmt.Print("请输入要修改的内容")
					fmt.Scan(&content)
					strings.TrimSpace(field)
					data[k][field] = content
				} else {
					fmt.Println("输入有误")
				}
			}
		}

		if !flag {
			fmt.Println("用户不存在")
		}
		fmt.Print("是否继续修改(y/n)")
		fmt.Scan(&txt)
		txt = strings.ToLower(strings.TrimSpace(txt))
		if txt == "y" {
			continue
		} else if txt == "n" {
			return
		} else {
			fmt.Println("输入有误从新输入(y/n)")
		}

	}
}
func deleteData() {
	var (
		name, txt string
		tmp       = make([]map[string]string, 0)
	)

	for {
		flag := false
		fmt.Print("请输入要删除的用户名称")
		fmt.Scan(&name)
		strings.TrimSpace(name)
		fmt.Println("正在验证用户信息")
		for k := range data {
			if data[k]["name"] == name {
				copy(data[k:], data[k+1:])
				tmp = data[:len(data)-1]
				fmt.Println("用户以删除")
				data = tmp
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("用户不存在")
		}
		fmt.Print("是否继续删除用户(y/n)")
		fmt.Scan(&txt)
		strings.TrimSpace(strings.ToLower(txt))
		if txt == "y" {
			continue
		} else {
			return
		}
	}
}

func
entry() {
	var count int
	var num int
	for count < 3 {
		fmt.Print(`
		用户管理
	1. 添加用户
	2. 查询用户信息
	3. 更新用户信息
	4. 删除用户
	5. 退出
`)
		fmt.Print("请输入序号：")
		if _, err := fmt.Scan(&num); err != nil {
			fmt.Println("请输入正确的范围数字, 1~5")
			count++
			continue
		}
		switch num {
		case 1:
			ret := addUser()
			if ret != nil {
				continue
			}
		case 2:
			selectData()
		case 3:
			updateData()
		case 4:
			deleteData()
		case 5:
			fmt.Println("-_-!!!!")
			return
		default:
			fmt.Println("请输入 1~5 范围的数字")
			count++
		}
	}

}

func
main() {
	entry()
}
