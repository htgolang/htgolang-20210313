package main

import "fmt"

/*
冒泡排序
 */
func main()  {
	var s []int = []int{2,3,12,90,22,1}
	fmt.Printf("调整前的值%#v\n",s)
	for i:=0;i<len(s);i++{
		for j:=i+1;j<len(s);j++ {
			if s[j] > s[i] {
				s[i],s[j] = s[j],s[i]
			}
		}
	}
	fmt.Printf("调整后的值%#v\n",s)
}
