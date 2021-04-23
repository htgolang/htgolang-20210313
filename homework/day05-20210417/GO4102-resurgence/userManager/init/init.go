package init

import (
	"fmt"
	"os"
	. "userManager/apps/todo"
	. "userManager/apps/user"
	. "userManager/config"
	. "userManager/utils"
)

func init() {
	if err := VerifyPWD(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	UserRegister("显示用户信息", GetUserList)
	UserRegister("添加用户信息", AddUser)
	UserRegister("更新用户信息", UpdateUser)
	UserRegister("删除用户信息", DeleteUser)
	UserRegister("退出", func(map[string]*User) {
		os.Exit(0)
	})

	TodoRegister("显示Todo信息", GetTodoList)
	TodoRegister("添加Todo信息", AddTodo)
	TodoRegister("更新Todo信息", UpdateTodo)
	TodoRegister("删除Todo信息", DeleteTodo)
	TodoRegister("退出", func(map[string]*Todo) {
		os.Exit(0)
	})
}
