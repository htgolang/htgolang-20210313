package user

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Addr string `json:"addr"`
	Tel string `json:"tel"`
	Pwd string `json:"pwd"`
}


//func Add() {
//	fmt.Println("add")
//	user.Add(user.Users)
//}

//func Modify() {
//	fmt.Println("modify")
//}
//
//func Delete() {
//	fmt.Println("delete")
//}
//
//func Query() {
//	fmt.Println("query")
//	user.FindUser(user.Users)
//}