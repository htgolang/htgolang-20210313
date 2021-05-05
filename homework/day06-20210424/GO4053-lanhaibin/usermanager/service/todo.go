package service

import (
	"errors"
	"usermanager/model"
)

type TodoService struct{}

func (TodoService) Get(id int) (*model.Todo, error) {
	todo, ok := model.Todos[id]
	if !ok {
		return nil, errors.New("任务不存在!")
	}
	return todo, nil
}

func (TodoService) Add(t model.Todo) error {
	err := ValidTodo(t)
	if err != nil {
		return err
	}
	id := model.TodoId
	t.Id = id
	model.Todos[id] = &t
	model.TodoId++
	return nil
}

func (TodoService) Delete(id int) error {
	var flag bool = false
	for _, todo := range model.Todos {
		if todo.Id == id {
			delete(model.Todos, id)
			flag = true
			break
		}
	}
	if flag {
		return nil
	}
	return errors.New("任务不存在!")
}

func (TodoService) Query(s string) model.TodoList {
	var qs model.TodoList
	for _, t := range model.Todos {
		if t.Contains(s) {
			qs = append(qs, t)
		}
	}
	return qs
}

func (TodoService) Modify(id int, t model.Todo) error {
	todo, ok := model.Todos[id]
	if !ok {
		return errors.New("任务不存在!")
	}
	t.Id = id
	err := ValidTodo(t)
	if err != nil {
		return err
	}
	todo.Name = t.Name
	todo.Priority = t.Priority
	todo.Desc = t.Desc
	todo.Progress = t.Progress
	todo.State = t.State
	todo.Principal = t.Principal
	todo.CreateAt = t.CreateAt
	todo.EndAt = t.EndAt
	todo.CompleteAt = t.CompleteAt
	return nil

}

func ValidTodo(t model.Todo) error {
	if len(t.Name) == 0 {
		return errors.New("任务名不能为空!")
	}

	us := UserService{}
	_, err := us.Get(t.Principal)
	if err != nil {
		return err
	}
	for _, todo := range model.Todos {
		if todo.Name == t.Name && todo.Id != t.Id {
			return errors.New("任务已存在!")
		}
	}
	return nil
}
