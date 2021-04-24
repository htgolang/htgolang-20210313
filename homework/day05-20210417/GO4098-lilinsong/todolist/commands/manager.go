package commands

import (
	"errors"
	"strconv"
	"time"
	"todolist/commands/command"
	"todolist/utils/ioutlis"
	"todolist/view"
)

type taskManager struct {
	tasks         map[string]*command.Tasks
	user          map[string]*command.Users
	command       map[string]*command.Command
	loginCallback command.LonginCallback
}

func newTaskManager() *taskManager {
	return &taskManager{
		tasks:   make(map[string]*command.Tasks),
		user:    make(map[string]*command.Users),
		command: map[string]*command.Command{},
	}
}

var mgr = newTaskManager()

func (t *taskManager) tasksRegister(task, priority, desc, schedule, status, name string, startTime, endTime *time.Time) {
	id := len(t.tasks) + 1
	t.tasks[task] = command.NewTasks(strconv.Itoa(id), task, priority, desc, schedule, status, name, startTime, endTime)
}

func TasksRegister(task, priority, desc, schedule, status, name string, startTime, endTime *time.Time) {
	mgr.tasksRegister(task, priority, desc, schedule, status, name, startTime, endTime)
}

func (t *taskManager) userRegister(name string, flag bool) {
	t.user[name] = command.NewUsers(name, flag)
}

func UserRegister(name string, flag bool) {
	mgr.userRegister(name, flag)
}

func (t *taskManager) loginCallbackRegister(callback command.LonginCallback) {
	t.loginCallback = callback
}

func LoginCallbackRegister(callback command.LonginCallback) {
	mgr.loginCallbackRegister(callback)
}

func (t taskManager) getUser(name string) (string, bool, error) {
	if _, ok := t.user[name]; ok {
		return name, t.user[name].Flag, nil
	}
	return "", false, errors.New("用户不存在")
}

func GetUser(name string) (string, bool, error) {
	return mgr.getUser(name)
}

func (t taskManager) loginPrompt() {
	view.ViewCmd.UserLogin()
}

func (t taskManager) menuPrompt() {
	view.ViewCmd.TaskMenu(t.command)
}

func (t *taskManager) commandRegister(task string, callback command.Callback) {
	t.command[strconv.Itoa(len(t.command)+1)] = command.NewCommand(task, callback)
}

func CommandRegister(task string, callback command.Callback) {
	mgr.commandRegister(task, callback)
}

func Del(task string) {
	delete(mgr.tasks, task)
}

func (t taskManager) run() {
	t.loginPrompt()
	if t.loginCallback != nil {
		if !t.loginCallback() {
			return
		}
	}
	for {
		t.menuPrompt()
		num := ioutlis.Input("请输入要选择的序号: ")
		if _, ok := t.command[num]; !ok {
			ioutlis.Error("输入的指令有误 ")
			continue
		}
		t.command[num].Callback()

	}
}

func GetTasks() map[string]*command.Tasks {
	return mgr.tasks
}

func Run() {
	BakFile()
	var json1 JsonPersist
	json1.Save()

	mgr.run()
}
