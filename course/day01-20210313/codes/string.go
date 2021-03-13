package main

import "fmt"

func main() {
	var desc string = "a\nb"
	var desc2 string = "a\\nb"
	var raw string = `
a\nb
fffds
fdsfds
fdsfds
`
	fmt.Println(desc)
	fmt.Println(desc2)
	fmt.Println(raw)

	desc3 := "我叫" + "kk" // utf8 1,2,3
	// 0,1,2(我),3,4,5(叫),6(k),7(k)
	fmt.Println(desc3)
	fmt.Println(desc3[1])                      //索引从0开始
	fmt.Printf("%T, %c\n", desc3[7], desc3[7]) //索引从0开始
	fmt.Println(desc3[0:5])                    // [start,end)

	fmt.Printf("%s\n", raw)
}
