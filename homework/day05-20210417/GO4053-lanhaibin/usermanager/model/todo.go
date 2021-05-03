package model

import (
	"fmt"
	"strings"
	"time"
)

var TodoId int = 1

type Todo struct {
	Id         int
	Name       string
	Priority   string
	Desc       string
	Progress   string
	State      int
	Principal  int
	CreateAt   *time.Time
	EndAt      *time.Time
	CompleteAt *time.Time
}

func (t Todo) String() string {
	return fmt.Sprintf("id: %d, 名称: %s, 描述: %s", t.Id, t.Name, t.Desc)
}

func (t Todo) Contains(s string) bool {
	if strings.Contains(t.Name, s) {
		return true
	}
	if strings.Contains(t.Desc, s) {
		return true
	}
	return false
}

var Todos map[int]*Todo = make(map[int]*Todo)

type TodoList []*Todo
