package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成一个100以内的随机数，方便测试
START:
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var randomsug int = r.Intn(100)
	fmt.Println(randomsug)
	// 设置猜数的变量
	var cli_sug int
	//设置猜测次数
	for max_nums := 1; max_nums <= 5; max_nums++ {
		fmt.Print("请输入你想中奖的神奇数字:")
		fmt.Scan(&cli_sug)
		if cli_sug > randomsug {
			fmt.Println("兄die,你猜的太大了")
		} else if cli_sug < randomsug {
			fmt.Println("来嘛，大胆点")
		} else {
			fmt.Printf("牛鼻哦！中奖了,是否继续呀(y/n)?")
			var is_continue string
			fmt.Scan(&is_continue)
			if is_continue == "y" {
				//如果继续，则重置计数器
				max_nums = 0
				goto START
			}
			goto END
		}
	}
END:
}

/*
indows PowerShell
版权所有 (C) Microsoft Corporation。保留所有权利。

尝试新的跨平台 PowerShell https://aka.ms/pscore6

PS D:\go_obj\day01> go run .\randomsug.go
29
请输入你想中奖的神奇数字:88
兄die,你猜的太大了
请输入你想中奖的神奇数字:29
牛鼻哦！中奖了
PS D:\go_obj\day01> go run .\randomsug.go
33
请输入你想中奖的神奇数字:33
牛鼻哦！中奖了
exit status 1
PS D:\go_obj\day01> go run .\randomsug.go
69
请输入你想中奖的神奇数字:1
来嘛，大胆点
请输入你想中奖的神奇数字:2
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:4
来嘛，大胆点
请输入你想中奖的神奇数字:69
牛鼻哦！中奖了
exit status 1
PS D:\go_obj\day01> go run .\randomsug.v2.go
5
来嘛，大胆点
请输入你想中奖的神奇数字:1
来嘛，大胆点
请输入你想中奖的神奇数字:1
来嘛，大胆点
请输入你想中奖的神奇数字:1
来嘛，大胆点
请输入你想中奖的神奇数字:5
牛鼻哦！中奖了,是否继续呀(y/n)?y
21
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:4
来嘛，大胆点
请输入你想中奖的神奇数字:7
兄die,你猜的太大了
请输入你想中奖的神奇数字:8
兄die,你猜的太大了
PS D:\go_obj\day01> go run .\randomsug.v2.go
13
请输入你想中奖的神奇数字:3
来嘛，大胆点
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:13
牛鼻哦！中奖了,是否继续呀(y/n)?y
73
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:73
兄die,你猜的太大了
PS D:\go_obj\day01> go run .\randomsug.v2.go
59
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
请输入你想中奖的神奇数字:3
来嘛，大胆点
牛鼻哦！中奖了,是否继续呀(y/n)?y
70
请输入你想中奖的神奇数字:k
牛鼻哦！中奖了,是否继续呀(y/n)?y
67
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:99
兄die,你猜的太大了
请输入你想中奖的神奇数字:40
来嘛，大胆点
请输入你想中奖的神奇数字:33
来嘛，大胆点
请输入你想中奖的神奇数字:67
兄die,你猜的太大了
PS D:\go_obj\day01> go run .\randomsug.v2.go
# command-line-arguments
.\randomsug.v2.go:8:1: syntax error: non-declaration statement outside function body
PS D:\go_obj\day01> go run .\randomsug.v2.go
60
请输入你想中奖的神奇数字:3
来嘛，大胆点
来嘛，大胆点
请输入你想中奖的神奇数字:3
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:60
牛鼻哦！中奖了,是否继续呀(y/n)?y
36
请输入你想中奖的神奇数字:36
来嘛，大胆点
请输入你想中奖的神奇数字:1
来嘛，大胆点
请输入你想中奖的神奇数字:1
来嘛，大胆点
请输入你想中奖的神奇数字:1
来嘛，大胆点
请输入你想中奖的神奇数字:1
来嘛，大胆点
# command-line-arguments
.\randomsug.v2.go:19:8: randomsug declared but not used
70
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:70
牛鼻哦！中奖了,是否继续呀(y/n)?y
请输入你想中奖的神奇数字:70
牛鼻哦！中奖了,是否继续呀(y/n)?n
请输入你想中奖的神奇数字:f
牛鼻哦！中奖了,是否继续呀(y/n)?请输入你想中奖的神奇数字:exit status 3221225786
PS D:\go_obj\day01> go run .\randomsug.v2.go
# command-line-arguments
.\randomsug.v2.go:36:1: label END defined and not used
PS D:\go_obj\day01> go run .\randomsug.v2.go
91
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:3
来嘛，大胆点
请输入你想中奖的神奇数字:91
牛鼻哦！中奖了,是否继续呀(y/n)?y
89
请输入你想中奖的神奇数字:80
来嘛，大胆点
请输入你想中奖的神奇数字:89
牛鼻哦！中奖了,是否继续呀(y/n)?y
32
请输入你想中奖的神奇数字:32
牛鼻哦！中奖了,是否继续呀(y/n)?y
73
请输入你想中奖的神奇数字:33
来嘛，大胆点
请输入你想中奖的神奇数字:44
来嘛，大胆点
请输入你想中奖的神奇数字:80
兄die,你猜的太大了
请输入你想中奖的神奇数字:33
来嘛，大胆点
请输入你想中奖的神奇数字:73
牛鼻哦！中奖了,是否继续呀(y/n)?

*/
