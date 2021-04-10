package math

import "fmt"

var Avar = "avar" // 可以
var bvar = "bvar" // 不可以

func Atest() { // 可以
	fmt.Println("Atest")
}

func btest() { //不可以
	fmt.Println("btest")
}

func GetBvar() string {
	main()
	return bvar
}

// 一般只将main函数定义在main包下
func main() {
	fmt.Println("xxx")
}
