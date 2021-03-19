package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
启动程序时生成一个随机数(target)让用户猜测guess(让用户输入数据)
猜测太大了 提示 太大了
猜测太小了 提示 太小了
想等  提示 猜对了 => 询问是否继续 -> 继续 重新生成随机数 不继续退出

最多猜测5测
 */

func main() {
	var times int
	var guess_num int
	var ack string
	rand.Seed(time.Now().UnixNano())
	correct_num := rand.Intn(100)
	fmt.Println(correct_num)

START:
	for {
		fmt.Printf("please input a int number(0~100)，有5次机会")
		fmt.Scan(&guess_num)
		if guess_num == correct_num {
			fmt.Printf("正确结果是%d的，你猜对了！\n", correct_num)
			fmt.Print("是否继续猜数字Y/y:")
			fmt.Scan(&ack)
			if ack == "Y" || ack == "y" {
				times = 0
				correct_num = rand.Intn(100)
				goto START
			} else if ack == "N" || ack == "n" {
				goto END
			} else {
				fmt.Println("输入非法，重新开始！")
				goto START
			}
		} else if guess_num > correct_num {
			times++
			fmt.Printf("猜大了！,还剩%d次机会!\n", 5-times)
		} else if guess_num < correct_num {
			times++
			fmt.Printf("猜小了！还剩%d次机会\n", 5-times)
		} else {
			fmt.Println("输入不对，请重新输入！机会重新开始")
			times = 0
		}
		if times == 5 {
			break
		}
	}
END:
	fmt.Println("游戏结束，Bye")
}
