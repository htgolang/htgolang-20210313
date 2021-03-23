package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//初始化随机数种子
	rand.Seed(time.Now().Unix())
	//生成随机数
	// fmt.Println(rand.Intn(100))
	a := rand.Intn(101)
	var b int
	var c string
	BEXIT:
	for i := 1; i < 6; i++ {
		fmt.Println("请猜数字:")
		fmt.Scan(&b)
		if b == a {
			fmt.Println("恭喜你猜对了！")
			fmt.Println("是否需要继续(y|n)?")
			fmt.Scan(&c)
			switch c {
			case "y":
				fmt.Scan(&b)
				a := rand.Intn(101)
				if b == a {
					fmt.Println("恭喜你猜对了！")
					break BEXIT
				} else if b > a {
					fmt.Println("你猜大了！")
				} else if b < a {
					fmt.Println("你猜小了！")
				}
			case "n":
				break BEXIT
			}
		} else if b > a {
			fmt.Println("你猜大了！")
		} else if b < a {
			fmt.Println("你猜小了！")
		}
		fmt.Println(i)
	}
}