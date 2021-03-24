package main

import "fmt"

/*
最大元素换到最后
 */
func main()  {
	var s []int = []int{9, 8, 19, 10, 2, 8}
	fmt.Printf("调整前的值%#v\n",s)
	/*根据index交换切片的值
	s[2],s[1] = s[1],s[2]
	fmt.Println(s)
	 */
	var max int
	var max_index int
	for i:=0;i<len(s);i ++ {
		if s[i] > max {
			max=s[i]
			max_index = i
		}
	}
	s[max_index],s[len(s)-1] = s[len(s)-1],s[max_index]
	fmt.Printf("调整后的值%#v\n",s)
}
