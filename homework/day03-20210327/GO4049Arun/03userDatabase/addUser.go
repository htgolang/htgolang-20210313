package main
import (
	"fmt"
	"regexp"
	"strconv"
)

func add(users []map[string]string) ([]map[string]string,error) {
	name := input(users, "请输入用户名(1~10非空字符,不能数字开头):", "name")
	addr := input(users, "请输入地址(2~20非空字符,不能数字开头):", "addr")
	phoneNum := input(users, "请输入联系方式(11位数字):", "phoneNum")
	users = append(users, map[string]string{
		"id":   strconv.Itoa(genId(users)),
		"name": name,
		"addr": addr,
		"tel":  phoneNum,
	})
	return users,nil
}

func usersPrint(users []map[string]string) {
	for i := 0; i < len(users); i++ {
		//golang map取key的时候是无序的,稍后解决该问题,临时指定了
		//go fmt.Printf中文对齐需要自己单独处理,稍后解决
		fmt.Printf("|id:%-2s|name:%-10s|tel:%-11s|addr:%-20s |\n",users[i]["id"],users[i]["name"],users[i]["tel"],users[i]["addr"])
	}
}

func genId(users []map[string]string) int {
	id := 0
	for _, user := range users {
		i, _ := strconv.Atoi(user["id"])
		if i > id {
			id = i
		}
	}
	return id + 1
}

func checkExistUser(users []map[string]string, name string) bool {
	for i := 0; i < len(users); i++ {
		if name == users[i]["name"] {
			return true
		}
	}
	return false
}

func checkName(alphabet string) bool {
	//一<\u4e00>~龥<\u9fa5>(基本汉字范围)
	if ok, _ := regexp.MatchString("^[a-zA-Z\u4e00-\u9fa5]\\S{1,9}$", alphabet); !ok {
		return false
	}
	return true
}

func checkAddr(alphabet string) bool {
	//一<\u4e00>~龥<\u9fa5>(基本汉字范围)
	if ok, _ := regexp.MatchString("^[a-zA-Z\u4e00-\u9fa5]{2,20}$", alphabet); !ok {
		return false
	}
	return true
}

func checkPhoneNumber(phone string) bool {
	if ok, _ := regexp.MatchString("^[0-9]{11}$", phone); !ok {
		return false
	}
	return true
}
func input(users []map[string]string, prompt string, flag string) string {
	var text string
	for {
		fmt.Print(prompt)
		fmt.Scan(&text)
		if "name" == flag {
			if checkName(text) {
				// 验证 用户名不能重复
				if checkExistUser(users, text) {
					fmt.Printf("用户名%s已存在,请重新输入!\n", text)
					continue
				}
				break
			} else {
				fmt.Println("名字输入不符合规范,请重新输入!")
			}
		} else if "phoneNum" == flag {
			if checkPhoneNumber(text) {
				break
			} else {
				fmt.Println("号码输入不符合规范,请重新输入!")
			}
		} else {
			if checkAddr(text) {
				break
			} else {
				fmt.Println("地址输入不符合规范,请重新输入!")
			}
		}
	}
	return text
}
