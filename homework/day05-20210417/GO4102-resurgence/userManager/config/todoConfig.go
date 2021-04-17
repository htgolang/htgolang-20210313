package config

import (
	"strconv"
	"time"
)

type Todo struct {
	id, Nice, Status             int
	Name, Desc, Plan, Head       string
	BeginTime, EndTime, OverTime time.Time
}

func NewTodo(Nice, Status int, Name, Desc, Plan, Head string, BeginTime, EndTime, OverTime time.Time) (*Todo, string) {
	Tid++
	return &Todo{
		id:        Tid,
		Nice:      Nice,
		Status:    Status,
		Name:      Name,
		Desc:      Desc,
		Plan:      Plan,
		Head:      Head,
		BeginTime: BeginTime,
		EndTime:   EndTime,
		OverTime:  OverTime,
	}, strconv.Itoa(Tid)
}

var (
	Tid         = 0
	TodoMap     = make(map[string]*Todo)
	TodoFuncMap = make(map[string]ft)
	TodoMenuMap = make(map[string]string)
)

type ft func(map[string]*Todo)

func TodoRegister(desc string, callback ft) {
	idx := strconv.Itoa(Tid)
	TodoMenuMap[idx] = desc
	TodoFuncMap[idx] = callback
	Tid++
}
