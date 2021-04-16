package consoles

import (
	"github.com/elliotchance/orderedmap"
)

var id int = 2

var ManagementView = orderedmap.NewOrderedMap()

func Register(value func()) {
	ManagementView.Set(id, value)
	id++
}

var Menu = map[int]string{
	1: "退出",
	2: "添加用户",
	3: "删除用户",
	4: "查询用户信息",
	5: "修改用户信息",
}


