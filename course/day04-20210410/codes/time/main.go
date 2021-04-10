package main

import (
	"fmt"
	"time"
)

func main() {
	// time
	// 时间: 字符串格式 unix时间戳 结构化格式
	now := time.Now()
	fmt.Printf("%T %v\n", now, now)
	fmt.Println(now.Unix())
	fmt.Println(now.Format("2006-01-02"))           //格式 年2006 月01 日02 时03/15 分04 秒05
	fmt.Println(now.Format("03:04:05"))             //格式 年2006 月01 日02 时03/15 分04 秒05
	fmt.Println(now.Format("15:04:05"))             //格式 年2006 月01 日02 时03/15 分04 秒05
	fmt.Println(now.Format("2006-01-02 15:04:05"))  //格式 年2006 月01 日02 时03/15 分04 秒05
	fmt.Println(now.Format("2006年01月02日15时04分05秒")) //格式 年2006 月01 日02 时03/15 分04 秒05

	// unixtime => Time
	fmt.Println(time.Unix(1618045552, 0))
	// 字符串 => Time
	fmt.Println(time.Parse("2006-01-02", "2021-05-10"))
	fmt.Println(time.Parse("15-01-02", "2021-05-10"))
	fmt.Println(time.Parse("2006-01-02 15:04:05", "2021-05-10 05:05:05"))
	// 时间区间
}
