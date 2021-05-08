package handles

import (
	"fmt"
)

var id int = 2

type Callback func()

type Tips struct {
	Id   int
	desc string
	Comm Callback
}

type tip []*Tips

var Tip = make(tip, 0)

//将函数自动注册到相应的slice
func Register(desc string, call Callback) {

	Tip = append(Tip, &Tips{id, desc, call})

	id++
}

//将提示信息打印出来
func Prompt() {
	//prompts := Prom
	fmt.Println("---请选择想要执行的操作---")
	fmt.Println("1，退出")

	for _, v := range Tip {
		fmt.Printf("%d, %s\n", v.Id, v.desc)
	}

}

func (*tip) GetID(id int) Callback {

	for _, value := range Tip {

		if id == value.Id {
			return value.Comm
		}
	}
	return nil
}
