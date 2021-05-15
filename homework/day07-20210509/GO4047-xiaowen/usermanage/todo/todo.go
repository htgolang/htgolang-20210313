package todo

import (
	"fmt"
	"strconv"
	"strings"
	// "time"
)

//增加一个todolist管理，涉及任务增删改查，任务字段：id，名称，优先级，说明，进度，状态，负责人，开始时间，结束时间，完成时间等属性



var TodoList = make(map[string]*TodoStr)

func TaskAdd() {
	TodoList[strconv.Itoa(Gid())] = OutputTodo()	
}


func TaskDel(){
	var id string
	fmt.Println("请输入要删除id:")
	fmt.Scan(&id)
	delete(TodoList, id)
	fmt.Printf("%v", TodoList)
}

func TaskMod(){
	var str,de string
	var id int
	fmt.Print("请输入要修改的id: ")
	fmt.Scan(&id)
	fmt.Println("请输入要修改的标题(taskname|head): ")
	fmt.Scan(&str)
	fmt.Println("请输入要修改的内容:")
	fmt.Scan(&de)
	if str == "taskname"{
		TodoList[strconv.Itoa(id)].TaskName = de
	} else if str == "head"{
		TodoList[strconv.Itoa(id)].Head = de
	} else{
		fmt.Println("输入有误")
		return
	}
	fmt.Printf("taskname: %v head: %v\n", TodoList[strconv.Itoa(id)].TaskName, TodoList[strconv.Itoa(id)].Head)
}

func TaskQuery(){
	var str string
	fmt.Println(TodoList)
	fmt.Println("请输入要查询的内容:")
	fmt.Scan(&str)
	for k,v := range TodoList{
		fmt.Println(k)
		// ks, _ := strconv.Atoi(k)
		if strings.Contains(v.TaskName, str) || strings.Contains(v.Head, str){
			fmt.Println(TodoList[k].Id, TodoList[k].TaskName, TodoList[k].Head)
	// 		fmt.Printf("id: %d  taskname: %v  priority: %v  describe: %v  progress: %0.2F  status: %v  head: %v  starttime: %v  endtime: %v  completionTime: %v\n", id, taskname, 
	// priority, describe, progress, status, head, startTime, endTime, completionTime)
			// fmt.Println("haha")
		}else{
			fmt.Println("不存在")
		}
	}
	// if strings.Contains(TodoList.TaskName, str) || strings.Contains(TodoList.Head, str){

	// }

}