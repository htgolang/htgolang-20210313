package tips

import (
	"strconv"
)

var taskid int = 6
var TaskMes [][2]string = [][2]string{}

func TaskPro(desc string){
	TaskMes = append(TaskMes, [2]string{strconv.Itoa(taskid), desc})
	taskid++
}