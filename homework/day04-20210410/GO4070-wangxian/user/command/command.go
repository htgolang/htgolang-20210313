package command

import "strconv"

var id = 2
var Commands = map[string]func(){}

// "2": usermanage.Add,
// "3": usermanage.Delete,
// "4": usermanage.Modify,
// "5": usermanage.Query}

var Prompt = [][2]string{}

// {"2", "添加"},
// {"3", "删除"},
// {"4", "修改"},
// {"5", "查询"}}

func Register(desc string, callback func()) {
	Commands[strconv.Itoa(id)] = callback
	Prompt = append(Prompt, [2]string{strconv.Itoa(id), desc})
	id++
}
