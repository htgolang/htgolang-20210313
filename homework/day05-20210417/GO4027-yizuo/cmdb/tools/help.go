package tools

import "fmt"

// 系统清单帮助信息
func SystemPrompt() {
	str := `--------------------------------------------------
	***********管理系统清单***********
	    user      ）用户管理系统
	    todolist  ）todolist管理
	    quit      ）退出系统
	********************************
	`
	fmt.Println(str)
}

// 用户管理系统帮助信息
func AdminHelpPrompt() {
	str := `--------------------------------------------------
	***********用户管理系统***********
	    add    ）添加用户
	    mod    ）修改用户
	    del    ）删除用户
	    query  ）搜索信息
	    list   ）打印用户列表
	    help   ) 帮助信息
	    quit   ）退出系统
	********************************
	`
	fmt.Println(str)
}

// todolist管理系统帮助信息
func ToDoListHelpPrompt() {
	str := `--------------------------------------------------
	***********todolist管理系统***********
	    add    ）添加数据
	    mod    ）修改数据
	    del    ）删除数据
	    query  ）搜索数据
	    list   ）打印数据列表
	    help   ) 帮助信息
	    quit   ）退出系统
	********************************
	`
	fmt.Println(str)
}
