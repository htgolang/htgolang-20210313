package todo

import (
	"fmt"
	. "userManager/config"
)

func GetTodoList(todoMap map[string]*Todo) {
	fmt.Println("get todo")
}

func AddTodo(todoMap map[string]*Todo) {
	fmt.Println("add todo")
}

func UpdateTodo(todoMap map[string]*Todo) {
	fmt.Println("update todo")
}

func DeleteTodo(todoMap map[string]*Todo) {
	fmt.Println("delete todo")
}
