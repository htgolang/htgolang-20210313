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
	for i := 1; i < 6; i++ {
		fmt.Println("请猜数字:")
		fmt.Scan(&b)
		if b == a {
			fmt.Println("恭喜你猜对了！")
			break
		} else if b > a {
			fmt.Println("你猜大了！")
		} else if b < a {
			fmt.Println("你猜小了！")
		}
		fmt.Println(i)
	}
}