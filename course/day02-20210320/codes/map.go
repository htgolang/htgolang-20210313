package main

import "fmt"

func main() {
	// name => cnt
	// 记录每个学生上课次数
	// key => string
	// value => int
	var stats map[string]int

	fmt.Printf("%T, %v, %v\n", stats, stats, stats == nil)

	// 赋值
	stats = map[string]int{}
	fmt.Printf("%v, %v\n", stats, stats == nil)

	stats = map[string]int{"黄鑫": 3, "kiss": 5}
	fmt.Printf("%v, %v\n", stats, stats == nil)

	stats = make(map[string]int)
	fmt.Printf("%v, %v\n", stats, stats == nil)

	stats = map[string]int{"黄鑫": 3, "kiss": 5}
	fmt.Println(len(stats))

	// 获取元素
	fmt.Println(stats["黄鑫"])
	fmt.Println(stats["黄鑫1"])

	var (
		v  int
		ok bool
	)

	v = stats["黄鑫"]
	fmt.Println(v)
	v = stats["黄鑫1"]
	fmt.Println(v)
	v, ok = stats["黄鑫"]
	fmt.Println(v, ok)
	v, ok = stats["黄鑫1"]
	fmt.Println(v, ok)
	// 增加元素
	stats["王辉"] = 5
	fmt.Println(stats)
	// 修改元素
	stats["王辉"] = 15
	fmt.Println(stats)
	//删除元素
	stats["xxxxx"] = 5
	fmt.Println(stats)
	delete(stats, "xxxxx")
	delete(stats, "xxxxx1")
	fmt.Println(stats)

	// 遍历
	// for range
	for key := range stats {
		fmt.Println(key, stats[key])
	}
	for k, v := range stats {
		fmt.Println(k, v)
	}
}
