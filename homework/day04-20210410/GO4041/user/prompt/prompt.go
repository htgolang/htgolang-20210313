package prompt

import "fmt"

func Prompt() {
	fmt.Println(`
欢迎使用用户管理系统
1. 添加
2. 修改
3. 删除
4. 查询
5. 退出
	`)
}

func Input(prompt string) string {
	fmt.Print(prompt)
	var txt string
	fmt.Scan(&txt)
	return txt
}
