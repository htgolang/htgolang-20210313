package command

import (
	"strconv"

)



var C map[string]func() = map[string]func(){}
var id int = 2
func Comm(f func()) {
	C[strconv.Itoa(id)] = f
	id++
}
// type UserDait struct{
// 	Useradd  func()
// 	Userquery func()
// 	Userdel func()
// 	Usermod func()
// }

// var id int = 2
// func Comm(f func()){
// 	U := UserDait
// 	U.
// }