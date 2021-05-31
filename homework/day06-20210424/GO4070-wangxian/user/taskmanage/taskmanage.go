package taskmanage

import (
	"fmt"
	"strconv"
	"strings"
	"user/module"
	"user/utils"
)

//生成taskid
func generateTaskID() int {
	id := 0
	for k := range module.Tasks {
		if k > id {
			id = k
		}
	}
	return id + 1
}

func TaskAdd() {
	tid := generateTaskID()
	newTask := utils.InputTask(tid)
	module.Tasks[tid] = newTask
	fmt.Println(module.Tasks)
}

func TaskDelete() {
	tid, err := strconv.Atoi(utils.Input("请输入要删除的任务Id:"))
	if err != nil {
		fmt.Println("输入错误")
		return
	}

	if v, ok := module.Tasks[tid]; ok {
		v.Print()
		confirm := utils.Input("是否确认删除(y/n)?")
		switch confirm {
		case "y", "Y":
			delete(module.Tasks, tid)
			fmt.Println(module.Tasks)
		case "n":
			fmt.Println("取消删除")
		default:
			fmt.Println("输入错误")
		}
	} else {
		fmt.Println("task不存在")
	}
}

func TaskModify() {
	tid, err := strconv.Atoi(utils.Input("请输入要删除的任务Id:"))
	if err != nil {
		fmt.Println("输入错误")
		return
	}

	if v, ok := module.Tasks[tid]; ok {
		v.Print()
		confirm := utils.Input("是否确认删除(y/n)?")
		switch confirm {
		case "y", "Y":
			tmpTask := utils.InputTask(tid)
			module.Tasks[tid] = tmpTask
			fmt.Println(module.Tasks)
		case "n":
			fmt.Println("取消修改")
		default:
			fmt.Println("输入错误")
		}
	} else {
		fmt.Println("task不存在")
	}
}

func TaskQuery() {
	keyWord := utils.Input("请输入要查询的关键字:")

	flag := false
	for _, v := range module.Tasks {
		//id不进行关键字匹配
		if strings.Contains(v.Name, keyWord) || strings.Contains(v.Priority, keyWord) ||strings.Contains(v.Desc, keyWord) || strings.Contains(v.Status, keyWord) || strings.Contains(v.Progress, keyWord) || strings.Contains(v.Owner, keyWord) || strings.Contains(v.StartAt.Format("2006-01-02"), keyWord) {
			v.Print()
			flag = true
		}
	}

	if !flag {
		fmt.Println("没有匹配任务")
	}

}
