package handles

import (
	"fmt"
	"strconv"
)

var id int =2

var Commands =map[string]func(){}

var Prom =[][2]string{}

func Prompt(){
	prompts := Prom
	fmt.Println("1、退出")

	for _,v :=range prompts{
		fmt.Printf("%s,%s\n",v[0],v[1])
	}

}

//将函数自动注册到相应的map和slice
func Register(desc string,callback func()){

	Commands[strconv.Itoa(id)]=callback
	Prom =append(Prom,[2]string{strconv.Itoa(id),desc})
	id ++
}