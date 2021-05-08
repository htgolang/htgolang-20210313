package init

import (
	"cmdb/handles"
	"cmdb/users"
)

func init(){

	handles.Register("添加",users.AddUser)
	handles.Register("删除",users.DeletUser)
	handles.Register("修改",users.ModifyUsers)
	handles.Register("查询",users.QueryUsers)

}
