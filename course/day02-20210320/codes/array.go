package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var nums [10]int
	fmt.Printf("%T, %v\n", nums, nums)
	// 零值: 由N个T类型的零值组成的数组

	var names [10]string
	fmt.Printf("%T, %q\n", names, names)

	// 字面量
	nums = [10]int{} //[0, 0,0,0,0,0,0,0,0,0]
	fmt.Println(nums)

	nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)

	nums = [10]int{1, 2, 3} // {1, 2, 3, 0,0,0,0,0,0,0}
	fmt.Println(nums)

	nums = [10]int{1: 1, 3: 2, 5: 3}
	fmt.Println(nums)

	// 语法糖 => go编译过程中进行转换 [length]T
	// 根据{}中元素推导length
	// [10]int
	nums = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)

	// [3]int
	// nums = [...]int{1, 2, 3} // {1, 2, 3, 0,0,0,0,0,0,0}
	// fmt.Println(nums)

	// fmt.Println([...]int{1, 2, 4: 5, 5: 5})

	// 内存: int 字节数量 * 元素的数量
	// 64位
	fmt.Println(unsafe.Sizeof(nums))

	// 操作
	// 算数运算 无
	// 关系运算 == !=
	nums2 := [10]int{1, 2}
	fmt.Println(nums == nums2)

	nums3 := [10]int{2, 1}
	fmt.Println(nums3 == nums2)
	// 逻辑运算，布尔运算
	// 函数
	// 长度
	fmt.Println(len(nums3))
	// 索引
	// 通过索引获取每个元素的值
	// 左->右:0,1,2,... ,len(array)-1
	// 无反向索引
	fmt.Println(nums[1], nums[9])

	// 遍历数组元素
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	// for range 遍历
	// 0 -> index
	// 1 -> 元素
	fmt.Println("range index:")
	for index := range nums {
		fmt.Println(index, nums[index])
	}

	fmt.Println("range index value:")
	for index, value := range nums {
		fmt.Println(index, value)
	}

	fmt.Println("_:")
	for _, value := range nums {
		fmt.Println(value)
	}

	// 修改元素
	fmt.Println(nums)
	nums[1] = 100
	fmt.Println(nums)
}
