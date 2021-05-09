package todolist

import (
	"fmt"
	"time"
)

/*
增加一个todolist管理，涉及任务增删改查，任务字段：id，名称，优先级，说明，进度，状态，负责人，开始时间，结束时间，完成时间等属性
*/

type Todo struct {
	Id         int
	Name       string
	Priority   string
	Desc       string
	Schedule   string
	Status     int
	Principal  string
	StartAt    time.Time
	EndAt      time.Time
	CompleteAt time.Time
}

type TodoList struct {
	Todos map[int]Todo
}

func TodoListAdd() {
	fmt.Println("TodoListAdd")
}

func TodoListDel() {
	fmt.Println("TodoListDel")
}

func TodoListMod() {
	fmt.Println("TodoListMod")
}

func TodoListQuery() {
	fmt.Println("TodoListQuery")
}

func TodoListListPrint() {
	fmt.Println("TodoListListPrint")
}
