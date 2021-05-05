package operation

import (
	"fmt"
	"usermanager/service"
	"usermanager/utils"
)

type DbOperation struct{}

func (DbOperation) ChangeDb() {
	fmt.Println("1. gob")
	fmt.Println("1. csv")
	type_ := utils.Input("请选择序列化方式:")
	switch type_ {
	case "1":
		service.Db.Type = "gob"
	case "2":
		service.Db.Type = "csv"
	default:
		fmt.Println("输入错误")
	}
}
