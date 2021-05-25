package operation

import (
	"fmt"
	"usermanager/service"
	"usermanager/utils"
)

type DbOperation struct{}

func (DbOperation) ChangeDb() {
	fmt.Println("1. gob")
	fmt.Println("2. csv")
	fmt.Println("3. json")
	type_ := utils.Input("请选择序列化方式:")
	switch type_ {
	case "1":
		service.Db.Type = "gob"
	case "2":
		service.Db.Type = "csv"
	case "3":
		service.Db.Type = "json"
	default:
		fmt.Println("输入错误")
	}
}
