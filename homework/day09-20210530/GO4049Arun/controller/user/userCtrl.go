package user
//package main

import (
	"encoding/gob"
	"mgr/model/user"
)
//var Users = []*user.User{
//	{1,"admin","北京市","16688889999","123"},
//	{2,"arun1","北京市","16899886688","123"},
//	{3,"Josh","America","15299886688","abc"},
//}
var Users []*user.User
var CurrentUser string
func init() {
	//fmt.Println(conf.Info.GobFilepath)
	//path:= "model/store/user.gob"
	gob.Register(&user.User{})

	//(1)Encode并写入文件
	//Create("gob").Write(Users)
	//Create("json").Write(users)
	//Create("csv").Write(users)

	//(2)解析读取
	//Users = ReadGob()
	Users = Create("json").Read()
	//Users = Create("csv").Read()
	//Users = Create("gob").Read()
	/*for _, u := range Users {
		fmt.Printf("Id:%v,Name:%v,Addr:%v,Tel:%v,Pwd:%v\n", u.Id,u.Name,u.Addr,u.Tel,u.Pwd)
	}*/
}