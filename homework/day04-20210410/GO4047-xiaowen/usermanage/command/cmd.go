package command

import (
	"strconv"

)



var C map[string]func(id int, args... map[string]map[string]string) = map[string]func(id int, args... map[string]map[string]string){}
var id int = 2
func Comm(f func(id int, args... map[string]map[string]string)) {
	C[strconv.Itoa(id)] = f
	id++
}