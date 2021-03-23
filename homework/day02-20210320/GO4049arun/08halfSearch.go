package main

import "fmt"

func main() {
	//Condition: slice is sorted from low to high by order
	var data = []int{2, 8, 9, 10, 19}
	var inputValue int
	fmt.Printf("请输入[2, 8, 9, 10, 19]中任意一个数:")
	fmt.Scan(&inputValue)
	var startIndex int = 0
	var maxIndex int = len(data) - 1
	for startIndex <= maxIndex {
		var midIndex int = (startIndex + maxIndex) / 2
		var midValue int = data[midIndex]
		if midValue == inputValue {
			fmt.Printf("存在,索引为:%d", midIndex)
			return
		} else if midValue > inputValue {
			maxIndex = midIndex - 1
		} else {
			startIndex = midIndex + 1
		}
	}
	fmt.Println("不存在,-1")
}

/*
PS E:\002golang\01NoteAll\20210320day02\02practice\Homework> go run .\08halfSearch.go
请输入[2, 8, 9, 10, 19]中任意一个数:2
存在,索引为:0
PS E:\002golang\01NoteAll\20210320day02\02practice\Homework> go run .\08halfSearch.go
请输入[2, 8, 9, 10, 19]中任意一个数:8
存在,索引为:1
PS E:\002golang\01NoteAll\20210320day02\02practice\Homework> go run .\08halfSearch.go
请输入[2, 8, 9, 10, 19]中任意一个数:9
存在,索引为:2
PS E:\002golang\01NoteAll\20210320day02\02practice\Homework> go run .\08halfSearch.go
请输入[2, 8, 9, 10, 19]中任意一个数:10
存在,索引为:3
PS E:\002golang\01NoteAll\20210320day02\02practice\Homework> go run .\08halfSearch.go
请输入[2, 8, 9, 10, 19]中任意一个数:19
存在,索引为:4
PS E:\002golang\01NoteAll\20210320day02\02practice\Homework> go run .\08halfSearch.go
请输入[2, 8, 9, 10, 19]中任意一个数:20
不存在,-1
*/
