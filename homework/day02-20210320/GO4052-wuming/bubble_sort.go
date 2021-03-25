
package main

import "fmt"

func main() {

	forSlice := []int{9,8,19,10,2,8}

for i :=0 ; i<len(forSlice)-1;i++{

 for j := 0 ; j< len(forSlice)-1-i;j++ {
	 if forSlice[j] > forSlice[j+1] {
		 forSlice[j+1], forSlice[j] = forSlice[j], forSlice[j+1]
	 }
 }

}

	fmt.Printf("调整后切片是 %v\n",forSlice)

}
