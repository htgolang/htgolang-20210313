package init

import (
	"user/command"
	"user/usermanage"
	"user/taskmanage"
)

func init() {

	command.Register(&command.UserCommands,"添加用户", usermanage.Add)
	command.Register(&command.UserCommands,"删除用户", usermanage.Delete)
	command.Register(&command.UserCommands,"修改用户", usermanage.Modify)
	command.Register(&command.UserCommands,"查询用户", usermanage.Query)

	command.Register(&command.TaskCommands, "添加任务", taskmanage.TaskAdd)
	command.Register(&command.TaskCommands, "删除任务", taskmanage.TaskDelete)
	command.Register(&command.TaskCommands, "修改任务", taskmanage.TaskModify)
	command.Register(&command.TaskCommands, "查询任务", taskmanage.TaskQuery)

	// command.Register(&command.TotalCommands, "用户管理", command.UserCommands.Prompt)
	// command.Register(&command.TotalCommands, "任务管理", command.TaskCommands.Prompt)
}

