package models

import "github.com/beego/beego/v2/adapter/orm"

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Asset))
	orm.RegisterModel(new(Role))
	orm.RegisterModel(new(Department))
	orm.RegisterModel(new(Module))
	orm.RegisterModel(new(RelRoleModule))
}
