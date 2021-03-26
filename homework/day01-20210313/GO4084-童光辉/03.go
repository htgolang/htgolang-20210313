package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//生成一个100以内的随机数，方便测试
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
			fmt.Println("牛鼻哦！中奖了")
			os.Exit(1)
		}
	}

}
