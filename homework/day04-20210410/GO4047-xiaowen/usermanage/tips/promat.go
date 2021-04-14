package tips

import (
	"strconv"
)

var id int = 2
var Mes [][2]string = [][2]string{}

func Promat(desc string){
	Mes = append(Mes, [2]string{strconv.Itoa(id), desc})
	id++
}