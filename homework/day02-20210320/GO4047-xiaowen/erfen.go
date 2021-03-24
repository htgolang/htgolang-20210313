package main
import (
	"fmt"
)

func main() {
	//初始将切片折半查找匹配，如果关键字小于索引值，则将当前索引设为右边界，再次和左边界对半查找（以此类推直至匹配成功），若是关键字大于索引值，则将索引设为左边界，和总长度对半查找（以此类推直至匹配成功）
	cap := []int{2, 8, 9, 10, 19}
	left := 0
	right := len(cap)
	mid := (left +right)/2
	var middle int
	fmt.Scan(&middle)
	for {
		if cap[mid] > middle {
			mid = (left + mid)/2
		} else if cap[mid] < middle {
			left = mid 
			mid = (left + right)/2
		} else  {
			fmt.Printf("索引为：%v", mid)
			return
		}
	}

}
