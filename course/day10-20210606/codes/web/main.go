package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const userFile = "users.json"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func readUsers() []User {
	users := make([]User, 0)
	fhandler, err := os.Open(userFile)
	if err != nil {
		return users
	}
	defer fhandler.Close()

	encoder := json.NewDecoder(fhandler)
	err = encoder.Decode(&users)
	if err != nil {
		return users
	}
	return users
}

func main() {
	// url => handler/handlerfunc => 先 生成html 再返回

	// 返回用户列表信息
	/*
		users.json中获取所用的用户信息 给浏览器返回
		<Table>
			<tr>
				<Td>1</td>
				<Td>Name</td>
			<tr>
		</Table>

		模板 fmt.Sprintf Printf Fprintf
		tpl 字符串 占位 + 数据 => 结果字符串
		模板技术: 定义模板
				由模板引擎 加载模板 通过指定数据 生成最终的字符串

				模板语法+使用引擎
		text/template
		html/template => 在web开发一定要用html => 字符转义(可以防止一些安全注入问题)
	*/

	addr := ":9999"
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		// 1. 麻烦
		// 2. 如果html结构写错了，排查问题麻烦，重新编译
		// 3. html页面 需求变化块
		users := readUsers()
		fmt.Fprint(w, "<!DOCTYPE html>")
		fmt.Fprint(w, "<html>")
		fmt.Fprint(w, "<head>")
		fmt.Fprint(w, "<meta charset='utf-8'>")
		fmt.Fprint(w, "<title>用户列表</title>")
		fmt.Fprint(w, "</head>")
		fmt.Fprint(w, "<body>")
		fmt.Fprint(w, "<table>")
		fmt.Fprint(w, "<thead>")
		fmt.Fprint(w, "<tr>")
		fmt.Fprint(w, "<th>ID</th>")
		fmt.Fprint(w, "<th>名字</th>")
		fmt.Fprint(w, "</tr>")
		fmt.Fprint(w, "</thead>")
		fmt.Fprint(w, "<tbody>")
		for _, user := range users {
			fmt.Fprint(w, "<tr>")
			fmt.Fprintf(w, "<td>%d</td>", user.ID)
			fmt.Fprintf(w, "<td>%s</td>", user.Name)
			fmt.Fprint(w, "</tr>")
		}
		fmt.Fprint(w, "</tbody>")
		fmt.Fprint(w, "</table>")
		fmt.Fprint(w, "</body>")
		fmt.Fprint(w, "</html>")
	})
	http.ListenAndServe(addr, nil)
}
