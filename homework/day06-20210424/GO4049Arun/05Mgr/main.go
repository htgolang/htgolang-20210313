package main

import (
	"mgr/view"
)
func main() {
	view.Login()

}
/*
E:\002golang\01NoteAll\20210424day06\03homework\GO4049Arun>go run .
请输入用户名:admin
请输入密码:
恭喜admin登录成功!
*******************************
 ====欢迎使用多云管理平台====
    1. 退出
    2. 新增用户
    3. 修改用户
    4. 删除用户)
    5. 查询用户信息
*******************************
请输入指令:2
请输入用户名(1~10非空字符,不能数字开头):admin
用户名admin已存在,请重新输入!
max
请输入地址(2~20非空字符,不能数字开头):中国北京
请输入联系方式(11位数字):13580809090
请输入密码(3位)：456
*******************************
 ====欢迎使用多云管理平台====
    1. 退出
    2. 新增用户
    3. 修改用户
    4. 删除用户)
    5. 查询用户信息
*******************************
请输入指令:5
请输入要模糊查询的字符串:a
id:1,name:admin,tel:16688889999,addr:北京市;
id:2,name:arun1,tel:16899886688,addr:北京市;
id:3,name:Josh,tel:15299886688,addr:America;
id:4,name:max,tel:13580809090,addr:中国北京
*/