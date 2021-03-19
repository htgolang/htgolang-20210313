package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a int         //输入的数字
	var b int         //猜测范围的最大值
	var c int         //生成的随机数
	var choose string //是否继续
	fmt.Println("欢迎使用猜数字模拟器：")
START:
	fmt.Println("请输入你要猜测的范围（输入最大值即可，默认从0到最大值）：")
	fmt.Scan(&b)
	rand.Seed(time.Now().Unix())
	c = rand.Intn(b + 1)
	//  fmt.Println(c)
	fmt.Println("随机数已生成,您有五次猜测的机会，祝您玩的愉快！下面游戏开始，输入您猜测范围内的数字即可（必须为整数否则会异常！）")
	for i := 1; i <= 5; i++ {
		fmt.Printf("请输入你第%d次猜测的数字：", i)
		fmt.Scan(&a)
		if a == c {
			fmt.Printf("恭喜您在第%d次猜测正确，你真棒！\n", i)
			break
		} else if a < c {
			fmt.Printf("猜测错误，您猜测的数值要比正确数值更小，你还有%d次机会，加油！\n", (5 - i))
		} else if a > c {
			fmt.Printf("猜测错误，您猜测的数值要比正确数值更大，你还有%d次机会，加油！\n", (5 - i))
		}
		if i == 5 {
			fmt.Printf("五次机会已用完，生成的随机数为%d  \n", c)
		}
	}

	fmt.Printf("你还要继续游戏吗？（y/n）")
	fmt.Scan(&choose)
	if choose == "y" || choose == "Y" {
		fmt.Println("你选择继续游戏，正在为您重新加载！")
		goto START
	} else {
		fmt.Println("您选择退出游戏，感谢您的使用！")
	}
}
