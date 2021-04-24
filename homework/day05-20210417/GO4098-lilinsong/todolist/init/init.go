package init

import (
	"time"
	"todolist/commands"
	"todolist/config"
	"todolist/controllers"
)

func init() {
	commands.LoginCallbackRegister(config.LoginAuth)
}

func init() {
	commands.UserRegister("candy", true)
	commands.UserRegister("tank", false)
}

func init() {

	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 48)
	commands.TasksRegister("上天", "100", "上天", "一半", "未完成", "蜡笔小新", &startTime, &endTime)
	commands.TasksRegister("入地", "100", "入地", "一半", "未完成", "蜡笔小新", &startTime, &endTime)
	commands.TasksRegister("玩", "100", "玩", "一半", "未完成", "蜡笔小新", &startTime, &endTime)
	commands.TasksRegister("哈哈", "100", "哈哈", "一半", "未完成", "蜡笔小新", &startTime, &endTime)
}

func init() {
	controllers.InitInterface()
	commands.CommandRegister("退出", controllers.Mt.Quit)
	commands.CommandRegister("添加任务", controllers.Mt.Add)
	commands.CommandRegister("查询任务", controllers.Mt.Query)
	commands.CommandRegister("删除任务", controllers.Mt.Del)
	commands.CommandRegister("修改任务", controllers.Mt.Modify)
}
