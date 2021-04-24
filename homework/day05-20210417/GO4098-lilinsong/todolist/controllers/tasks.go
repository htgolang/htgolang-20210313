package controllers

import (
	"fmt"
	"os"
	"time"
	"todolist/commands"
	"todolist/commands/command"
	"todolist/config"
	"todolist/utils/ioutlis"
)

type managerTools struct {
	name string
}

var Mt mTools

func InitInterface() {
	Mt = managerTools{}
}

type mTools interface {
	Quit()
	Add()
	Query()
	Del()
	Modify()
}

var jp = commands.InitJson()

func (m managerTools) Quit() {
	os.Exit(0)
}

func (m managerTools) Add() {
	task := ioutlis.Input("请输入任务名称: ")
	if _, err := m.isTaskExists(task); err != nil {
		fmt.Println(err)
		return
	}
	pri := ioutlis.Input("任务优先级: ")
	desc := ioutlis.Input("任务描述: ")
	name := ioutlis.Input("用户名称: ")
	fmt.Println("输入时间格式 YYYY-mm-dd")
	tm := ioutlis.Input("完成时间: ")
	endTime, _ := time.Parse(config.TimeLayout, tm)
	startTime := time.Now()
	schedule := "刚开始"
	status := "未完成"
	commands.TasksRegister(task, pri, desc, schedule, status, name, &startTime, &endTime)

	jp.Save()
}

func (m managerTools) Del() {
	task := ioutlis.Input("请输入要删除的任务的名称: ")
	_, err := m.isTaskExists(task)
	if err != nil {
		fmt.Println(err)
		return
	}
	commands.Del(task)
}

func (m managerTools) Modify() {
	fmt.Println("修改")
}

func (m managerTools) Query() {
	task := ioutlis.Input("请输入要查询任务的名称: ")
	tasks, err := m.isTaskExists(task)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Id: %s,task: %s,user: %s,startTime: %s, status: %s\n",
		tasks[task].Id, tasks[task].Task, tasks[task].Name, tasks[task].StartTime.Format(config.TimeLayout), tasks[task].Status)
}

func (m managerTools) isTaskExists(task string) (map[string]*command.Tasks, error) {
	tasks := commands.GetTasks()
	if _, ok := tasks[task]; !ok {
		return nil, fmt.Errorf("任务不存在: ")
	}

	return tasks, nil
}
