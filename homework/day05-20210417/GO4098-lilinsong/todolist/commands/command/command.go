package command

import (
	"time"
)

type Callback func()

type LonginCallback func() bool

type Tasks struct {
	Id        string     `json:"id"`
	Task      string     `json:"task"`
	Priority  string     `json:"priority"`
	Desc      string     `json:"desc"`
	Schedule  string     `json:"schedule"`
	Status    string     `json:"status_name"`
	Name      string     `json:"name"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	OverTime  *time.Time `json:"over_time gender,omitempty"`
}

func NewTasks(id, task, priority, desc, schedule, status, name string, startTime, endTime *time.Time) *Tasks {
	return &Tasks{
		Id:        id,
		Task:      task,
		Priority:  priority,
		Desc:      desc,
		Schedule:  schedule,
		Status:    status,
		Name:      name,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

type Users struct {
	Name string
	Flag bool
}

func NewUsers(name string, flag bool) *Users {
	return &Users{
		Name: name,
		Flag: flag,
	}
}

type Command struct {
	Task     string
	Callback Callback
}

func NewCommand(task string, callback Callback) *Command {
	return &Command{
		Task:     task,
		Callback: callback,
	}
}
