package main

import "fmt"

// 将layer个盘子从t1 移动到 t3 借助 t2
func tower(t1, t2, t3 string, layer int) {
	if layer == 1 {
		fmt.Printf("%s -> %s\n", t1, t3)
		return
	}
	// 先将layer-1个从t1移动到t2 借助 t3
	tower(t1, t3, t2, layer-1)
	// t1 -> t3
	fmt.Printf("%s -> %s\n", t1, t3)

	// 将layer-1个t2 移动到t3 借助 t1
	tower(t2, t1, t3, layer-1)
}

func main() {
	tower("A", "B", "C", 4)
}
