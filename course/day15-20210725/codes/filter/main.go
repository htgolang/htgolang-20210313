package main

import (
	"fmt"

	"github.com/beego/beego/v2/server/web/context"

	"github.com/beego/beego/v2/server/web"
)

func main() {

	// BeforeStatic
	// session初始化
	// BeforeRouter
	// BeforExec
	// AfterExec
	// FinshRouter

	// 获取session 判断用户是否登陆
	// 需要在哪个pos进行处理
	// BeforeStatic 不能 因为所有的ssession都没有读取
	// session初始化
	// BeforeRouter 可以
	// BeforExec 可以
	// Prepare 可以
	// AfterExec 不可以 Action 已执行
	// FinshRouter 不可以 Action 已执行

	// /login 过滤未登录能访问的URL
	web.InsertFilter("/*", web.BeforeRouter, func(c *context.Context) {
		fmt.Println("filter")
		c.WriteString("filter")
		// 检查是否为未登录可以访问的URL
		// 是 return
		// 不是 检查session是用户是否存在
		// c.Input.Session
		// 存在 return
		// 不存在跳转 /login/
	}, web.WithReturnOnOutput(false))

	// 统计 每5分钟 请求次数
	// int[12] -
	// [0:00 - 0:05) 0
	// [0:05 - 0:10) 1
	// 0 - 3600

	// Filter
	// 获取当前时间 current % 3600 / 300 => index

	web.Get("/index/", func(c *context.Context) {
		fmt.Println("index")
		c.WriteString("index")
	})

	web.Run()
}
