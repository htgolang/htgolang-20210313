package model

import (
	"fmt"
	"strings"
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

func NewTodo(id int, name, priority, desc, schedule string, status int, principal string, startAt, endAt, completeAt time.Time) *Todo {
	return &Todo{
		Id:         id,
		Name:       name,
		Priority:   priority,
		Desc:       desc,
		Schedule:   schedule,
		Status:     status,
		Principal:  principal,
		StartAt:    startAt,
		EndAt:      endAt,
		CompleteAt: completeAt,
	}
}

func (t Todo) String() string {
	fmt.Println(strings.Repeat("=", 20))
	return fmt.Sprintf("任务id: %d\n任务名 %v\n优先级: %v\n说明: %v\n进度: %v\n状态: %v\n负责人: %v\n开始时间: %v\n结束时间: %v\n完成时间: %v\n",
		t.Id, t.Name, t.Priority, t.Desc, t.Schedule, t.Status, t.Principal, t.StartAt.Format("2006-01-02"), t.EndAt.Format("2006-01-02"), t.CompleteAt.Format("2006-01-02"))
}
