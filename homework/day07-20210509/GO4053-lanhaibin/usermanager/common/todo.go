package common

import (
	"usermanager/model"
	"usermanager/utils"
)

func InputTodo() model.Todo {
	todo := model.Todo{}
	todo.Name = utils.Input("请输入任务名:")
	todo.Priority = utils.Input("请输入优先级:")
	todo.Desc = utils.Input("请输入描述:")
	todo.Progress = utils.Input("请输入进度:")
	todo.State = utils.InputInt("请输入状态:")
	todo.Principal = utils.InputInt("请输入负责人id:")
	createAt := utils.InputTime("请输入开始时间:")
	endAt := utils.InputTime("请输入结束时间:")
	completeAt := utils.InputTime("请输入完成时间:")
	todo.CreateAt = &createAt
	todo.EndAt = &endAt
	todo.CompleteAt = &completeAt

	// us := service.UserService{}
	// _, err := us.Get(todo.Principal)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return model.Todo{}, errors.New("负责人不存在!")
	// }
	return todo
}
