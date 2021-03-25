package main

import "fmt"

/*  位置互换
func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	num := 0
	index := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > num {
			num = slice[i]
			index = i
		}
	}
	tmp := slice[len(slice)-1]
	slice[len(slice)-1] = num
	slice[index] = tmp
	//	fmt.Println(slice)
}
*/
//顺序不变
func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	num := 0
	index := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > num {
			num = slice[i]
			index = i
		}
	}
	tmp := append(slice[:index], slice[index+1:]...)
	slice = append(tmp, num)
	fmt.Println(slice)
}
