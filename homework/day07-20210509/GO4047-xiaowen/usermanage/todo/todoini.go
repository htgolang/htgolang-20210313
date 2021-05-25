package todo

import (
	"fmt"
	// "std/strings"
	"strconv"
	"strings"
	"time"
)

//增加一个todolist管理，涉及任务增删改查，任务字段：id，名称，优先级，说明，进度，状态，负责人，开始时间，结束时间，完成时间等属性



var id int = 1
var Mes [][2]string = [][2]string{}


type TodoStr struct{
	Id int
	TaskName string
	Priority int
	Describe string
	Progress float64
	Status  string
	Head  string
	StartTime time.Time
	EndTime  time.Time
	CompletionTime time.Time
}


// func TodoView(desc string){
// 	Mes = append(Mes, [2]string{strconv.Itoa(id), desc})
// 	id++
// }


func inputs(p string) string{
	var str string
	fmt.Println(p)
	fmt.Scan(&str)
	return strings.TrimSpace(str)
}

func OutputTodo() *TodoStr{
	id := Gid()
	taskname := inputs("请输入任务名称(string):")
	priority,_ := strconv.Atoi(inputs("请输入优先级(int):"))
	describe := inputs("请输入描述信息(string):")
	progress,_ := strconv.ParseFloat(inputs("请输入进度(float):"), 64)
	status := inputs("请输入状态(string):")
	head := inputs("请输入负责人(string):")

	return Res(id,priority,taskname,describe,status,head,progress,time.Now(),time.Now(),time.Now(),)
	
}



func Res(id, priority int, taskname,describe, status, head string, progress float64, startTime,endTime,completionTime time.Time) *TodoStr{
	fmt.Printf("id: %d  taskname: %v  priority: %v  describe: %v  progress: %0.2F  status: %v  head: %v  starttime: %v  endtime: %v  completionTime: %v\n", id, taskname, 
	priority, describe, progress, status, head, startTime, endTime, completionTime)
	return &TodoStr{
		id,
		taskname,
		priority,
		describe,
		progress,
		status,
		head,
		startTime,
		endTime,
		completionTime,
	}
}

func Gid() int{
	id := 0
	for k ,_ := range TodoList{
		i,_ := strconv.Atoi(k)
		if i > id{
			id = i
		}
	}
	return id + 1
}