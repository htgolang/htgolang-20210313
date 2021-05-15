package main

import (
	"fmt"
	cuser "mgr/controller/user"
)

//func TestCreate(t *testing.T) {
func main() {
	//var Users = []*user.User{
	//	{1,"admin","北京市","16688889999","123"},
	//	{2,"arun1","北京市","16899886688","123"},
	//	{3,"Josh","America","15299886688","abc"},
	//}
	//cuser.Create("gob").Write(Users)
	//cuser.Create("json").Write(Users)
	//cuser.Create("csv").Write(Users)


	//Users := cuser.Create("json").Read()
	//Users :=cuser.Create("csv").Read()
	Users := cuser.Create("gob").Read()
	for _, u := range Users {
		fmt.Printf("Id:%v,Name:%v,Addr:%v,Tel:%v,Pwd:%v\n", u.Id,u.Name,u.Addr,u.Tel,u.Pwd)
	}
}