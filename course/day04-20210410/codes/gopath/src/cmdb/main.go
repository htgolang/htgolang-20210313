package main

// 导入包
import (
	"cmdb/idc"
	"cmdb/user" // gopath src后面的目录路径
	"prometheus/agent"

	"cmdb/idc/xxxx"
)

func main() {

	// 在不一个包下面调用不同文件中定义的函数 => 通过包名.函数名称
	//
	user.Add()
	user.Modify()
	user.Delete()
	user.Find()

	// 调用idc 中的add
	idc.Add()

	// 调用prometheus agent
	agent.Add()

	xxxx.Add()
}
