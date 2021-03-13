package main

import "fmt"

var version string = "1.1.1"

func main() { // main作用域

	// 作用域：变量的可使用范围
	// {}显示的声明一个作用域范围

	var funcVersion string = "2.2.2"

	fmt.Println("version.1", version) //1.1.1
	// 可以
	var version string = "x.x.x.x"
	fmt.Println("version.2", version) // x.x.x.x

	// var funcVersion string = "yyyy"

	{ //A
		fmt.Println("version.3", version) // x.x.x.x
		version = "yyyyy"                 // main.version
		fmt.Println("version.4", version) // yyyyy
	}

	fmt.Println("version.5", version) // yyyyy

	{ //B
		fmt.Println("version.6", version) // main.version yyyy
		var version string = "zzz"        //???? 在B作用域内定义了version ???
		fmt.Println("version.7", version) // zzz
	}

	fmt.Println(version, funcVersion) // yyyy
}
