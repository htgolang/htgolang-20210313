package view

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
	"todolist/commands/command"
)

type viewManager struct {
}

func (v viewManager) TaskMenu(task map[string]*command.Command) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"序号", "说明"})
	for i := 1; i <= len(task); i++ {
		num := strconv.Itoa(i)
		table.Append([]string{
			num,
			task[num].Task,
		})
	}
	table.Render()
}

func (v viewManager) UserLogin() {
	fmt.Println(strings.Repeat("~", 45))
	fmt.Println("		任务管理系统")
	fmt.Println(strings.Repeat("~", 45))
	fmt.Println("用户登录")
}

var ViewCmd viewManager
