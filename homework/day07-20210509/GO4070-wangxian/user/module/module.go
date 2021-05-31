package module

import (
	"fmt"
	"time"
)

type User struct {
	Id     int
	Name   string
	Addr   string
	Tel    string
	Passwd string
	// Logged bool
}

func (u User) Print()  {
	fmt.Printf("id:%d, name: %s, addr: %s, tel: %s\n", u.Id, u.Name, u.Addr, u.Tel)
}

type Task struct {
	Id       int //创建时自动生成，不能修改
	Name     string
	Priority string
	Desc     string
	Status   string
	Progress string //应该是自动生成的，目前手动指定
	Owner    string
	StartAt  time.Time  //任务开始执行时自动生成，不能修改
	FinishAt *time.Time //任务完成时自动生成，不能修改
}

func (t Task) Print() {
	fmt.Printf("ID:%d, 任务名称:%s, 任务优先级:%s, 任务描述:%s, 任务状态:%s, 任务进度:%s, 负责人:%s, 开始时间:%s, 结束时间:%s\n", t.Id, t.Name, t.Priority, t.Desc, t.Status, t.Progress, t.Owner, t.StartAt.Format("2006-01-02"), t.FinishAt)
}

// var UserInfo = []User{
// 	{Id: 0, Name: "wx", Addr: "北京", Tel: "1348690", Passwd: "a3dcb8aae0d84122a1778e26da285857"},
// 	{Id: 3, Name: "jack", Addr: "成都", Tel: "15274590", Passwd: "a3dcb8aae0d84122a1778e26da285857"},
// 	{Id: 4, Name: "jim", Addr: "西安", Tel: "1528690", Passwd: "a3dcb8aae0d84122a1778e26da285857"},
// }

var Tasks = map[int]Task{
	0: {0, "服务器上架", "middle", "服务器上架", "Running", "%50", "wx", time.Now(), nil},
	1: {1, "代码发布", "lowest", "测试环境代码发布", "Failed", "%100", "wx", time.Now(), nil},
}
