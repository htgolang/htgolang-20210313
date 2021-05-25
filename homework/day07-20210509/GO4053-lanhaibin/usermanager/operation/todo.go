package operation

import (
	"fmt"
	"strings"
	"usermanager/common"
	"usermanager/service"
	"usermanager/utils"
)

type TodoOperation struct{}

func (TodoOperation) Add() {
	todo := common.InputTodo()
	ts := service.TodoService{}
	err := ts.Add(todo)
	if err != nil {
		fmt.Println(err)
	}
}

func (TodoOperation) Delete() {
	id := utils.InputInt("请输入要删除的任务id:")
	ts := service.TodoService{}
	todo, err := ts.Get(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("请确认任务信息:")
	fmt.Println(todo)
	ask := utils.Input("是否删除?(y/n)")
	if strings.ToLower(ask) != "y" && strings.ToLower(ask) != "yes" {
		return
	}
	err = ts.Delete(id)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (TodoOperation) Query() {
	s := utils.Input("请输入查询字符串:")
	ts := service.TodoService{}
	tl := ts.Query(s)
	if len(tl) != 0 {
		fmt.Println("查询到以下任务信息:")
		for _, todo := range tl {
			fmt.Println(todo)
		}
	} else {
		fmt.Println("未查询到任何任务!")
	}
}

func (TodoOperation) Modify() {
	id := utils.InputInt("请输入要修改的任务id:")
	ts := service.TodoService{}
	todo, err := ts.Get(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(todo)
	ask := utils.Input("是否修改?")
	if strings.ToLower(ask) != "y" && strings.ToLower(ask) != "yes" {
		return
	}
	t := common.InputTodo()
	fmt.Println(t)
	ask = utils.Input("是否修改?")
	if strings.ToLower(ask) != "y" && strings.ToLower(ask) != "yes" {
		fmt.Println(ask)
		return
	}
	err = ts.Modify(id, t)
	if err != nil {
		fmt.Println(err)
	}
}
