package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todolist/commands/command"
)

type Persist interface {
	Save()
}

type JsonPersist struct {
}

func (j JsonPersist) Save() {
	pwd, _ := os.Getwd()
	fileName := filepath.Join(pwd, "data/tasks.json")
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()


	encoder := json.NewEncoder(file)
	if err := encoder.Encode(GetTasks()); err != nil {
		fmt.Println(err)
		return
	}
}

func JsonLoad(name string) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Printf("文件读取失败:%v", err)
		return
	}
	decoder := json.NewDecoder(file)
	var tasks = map[string]*command.Tasks{}
	if err := decoder.Decode(&tasks); err != nil {
		fmt.Printf("json解析失败:%v", err)
		return
	}
	fmt.Println(tasks)
}

var JsonSave Persist

func InitJson() Persist {
	JsonSave = JsonPersist{}
	return JsonSave
}
