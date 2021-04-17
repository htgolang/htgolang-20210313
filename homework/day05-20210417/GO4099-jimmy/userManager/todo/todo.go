package todo

import (
	"fmt"
	"strconv"
	"strings"
	"userManager/common"
	"userManager/model"
)

var todoList = model.TodoList{Todos: make(map[int]model.Todo)}

func Add() {
	id := generateId()
	todo := common.InputTodo(id)
	todoList.Todos[id] = *todo
	fmt.Println("任务添加成功")
	fmt.Println(todo)
}
// id int, name, priority, desc, schedule string, status int, principal string, startAt, endAt, completeAt
func Query() {
	qc := common.Input("请输入要查询的字符串: ")
	for id, u := range todoList.Todos {
		if strings.Contains(u.Name, qc) || strings.Contains(u.Priority, qc) || strings.Contains(u.Desc, qc) ||
			strings.Contains(strconv.Itoa(id), qc) || strings.Contains(u.Schedule, qc) || strings.Contains(strconv.Itoa(u.Status), qc) ||
			strings.Contains(u.Principal, qc) || strings.Contains(u.StartAt.Format("2006-01-02"), qc) || strings.Contains(u.EndAt.Format("2006-01-02"), qc) ||
			strings.Contains(u.CompleteAt.Format("2006-01-02"), qc) {
			fmt.Println(u)
		}
	}
	fmt.Println(strings.Repeat("*", 20))
}

func Delete() {
	id, _ := strconv.Atoi(common.Input("请输入要删除的任务id: "))
	if u, ok := todoList.Todos[id]; ok {
		fmt.Println(u)
		switch common.Input("确认是否删除y/yes: ") {
		case "y", "yes":
			delete(todoList.Todos, id)
			fmt.Println("任务删除成功")
		default:
			fmt.Println("未执行删除动作")
		}
	} else {
		fmt.Println("任务未找到")
	}
}

func Modify() {
	id, _ := strconv.Atoi(common.Input("请输入要修改的任务id: "))
	if u, ok := todoList.Todos[id]; ok {
		fmt.Println(u)
		switch common.Input("确认是否编辑y/yes: ") {
		case "y", "yes":
			todoList.Todos[id] = *common.InputTodo(id)
			fmt.Println("成功修改任务信息")
			fmt.Println(todoList.Todos[id])
		default:
			fmt.Println("未执行修改动作")
		}
	} else {
		fmt.Println("任务未找到")
	}
}

func generateId() int {
	maxId := 0
	for k := range todoList.Todos {
		if k > maxId {
			maxId = k
		}
	}
	return maxId + 1
}
