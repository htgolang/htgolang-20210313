/*猜数字游戏 启动程序时生成一个随机数(target)让用户猜测guess(让用户输入数据) 猜测太大了 提示 太大了 猜测太小了 提示 太小了 想等 提示 猜对了 => 程序退出,
最多猜测5测*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //设置一个随机数种子,time.Now().UnixNano是当前操作系统的毫秒值
	num2 := rand.Intn(10)            //定义一个变量num2，并将一个0-99的随机数赋值给它
	var num1 int                     //定义一个变量num1，为整数型。
	fmt.Print("请输入你所猜的(0~10)数字:")
	fmt.Scan(&num1)
	fmt.Println(num1)
	for n := 1; n < 5; n++ {
		if num1 > num2 {
			fmt.Println("你猜的数字偏大了。")
			fmt.Println("你还有", 5-n, "次机会")
			fmt.Scan(&num1)
		} else if num1 < num2 {
			fmt.Println("你猜的数字偏小了。")
			fmt.Println("你还有", 5-n, "次机会")
			fmt.Scan(&num1)
		} else if num1 == num2 {
			fmt.Println("恭喜你猜对了。")
			break
		}
	}
}
