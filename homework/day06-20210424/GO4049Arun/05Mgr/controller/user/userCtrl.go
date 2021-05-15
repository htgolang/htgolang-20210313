package user
//package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"mgr/model/user"
	"mgr/conf"
	"mgr/model/log"
)

func writeGob(users []*user.User)  {
	absPath , _:= filepath.Abs(conf.Info.GobFilepath)
	file, _ := os.Create(absPath)
	defer file.Close()
	//(1)写入
	encoder := gob.NewEncoder(file)
	err := encoder.Encode(users)
	if err==nil{
		fmt.Println("writeGob初始化用户写入成功!")
	}
}

func ModifyGob(users []*user.User)  {
	absPath , _:= filepath.Abs(conf.Info.GobFilepath)
	//_err := os.Remove(absPath)
	//if _err!=nil{
	//	fmt.Println("delete failed!",_err)
	//}else {
	//	fmt.Println("delete success!")
	//}
	//好像用open不行,看看后续其它优化
	file, _ := os.Create(absPath)
	//(1)写入
	encoder := gob.NewEncoder(file)
	//fmt.Printf("%#v\n",users)
	err := encoder.Encode(users)
	file.Close()
	if err==nil{
		//fmt.Println("写入成功")
	}else {
		log.LogRecord("Gob文件写入失败,请检查原因")
	}
}

func ReadGob() []*user.User {
	absPath , _:= filepath.Abs(conf.Info.GobFilepath)
	fileDecode, _ := os.Open(absPath)
	defer fileDecode.Close()
	decoder := gob.NewDecoder(fileDecode)
	users2 := make([]*user.User, 0)
	//Decode存储到users2中
	err := decoder.Decode(&users2)
	if err == nil{
		//fmt.Println(err, users2)
		return users2
	}else{
		log.LogRecord("Parse err!")
		//fmt.Println("Parse err!")
		return nil
	}
	/*for _, u := range users2 {
		fmt.Printf("Id:%v,Name:%v,Addr:%v,Tel:%v,Pwd:%v\n", u.Id,u.Name,u.Addr,u.Tel,u.Pwd)
	}*/
}
var Users = []*user.User{
	{1,"admin","北京市","16688889999","123"},
	{2,"arun1","北京市","16899886688","123"},
	{3,"Josh","America","15299886688","abc"},
}
var CurrentUser string
func init() {
	//fmt.Println(conf.Info.GobFilepath)
	//path:= "model/store/user.gob"
	gob.Register(&user.User{})

	//(1)Encode并写入文件
	//ModifyGob(Users)
	//(2)解析读取
	Users = ReadGob()
	/*for _, u := range Users {
		fmt.Printf("Id:%v,Name:%v,Addr:%v,Tel:%v,Pwd:%v\n", u.Id,u.Name,u.Addr,u.Tel,u.Pwd)
	}*/
}