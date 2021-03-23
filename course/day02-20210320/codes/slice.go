package main

import "fmt"

func main() {
	var nums []int // var nums []int = nil
	// nil切片
	fmt.Printf("%T, %v, %v\n", nums, nums, nums == nil)

	// 内存 暂时不关注

	// 字面量
	nums = []int{} // 空切片
	fmt.Printf("%v, %v\n", nums, nums == nil)

	nums = []int{1, 2, 3, 4}
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)

	nums = []int{1: 1, 10: 2}
	fmt.Println(nums)

	// make([]int, length)
	// length 切片元素的数量
	nums = make([]int, 5)
	// 赋值一个有5个元素的int类型零值组成的切片
	fmt.Println(nums)
	/*
		make([]int, length, cap)
		length 元素数量
		cap 容量 => 底层数组的长度
		slice底层 => 数组存储 数组长度固定,
			当length == cap 再添加元素重新申请内存，拷贝，是否旧数组内存
		长度增长(元素数量) 增长 => ??
			=> 重新生成一个新可以容纳更多元素的数组（拷贝原来数组中的元素）
		底层数组的原长度:10, 增加一个元素,
			增长：10, 底层数组的现在长度20, 元素: 11, 9没有用
	*/
	nums = make([]int, 3, 10)
	fmt.Println(nums)

	// 切片操作 => 数组(v)，字符串，切片
	//
	numsarray := [5]int{1, 3, 5, 7, 8}
	fmt.Printf("%T, %v\n", numsarray[1:3], numsarray[1:3])
	// [start:end] start <= end <= length

	nums = numsarray[1:5]
	fmt.Println(nums) //? length = end -start, cap
	// 容量

	// 切片的切片操作
	fmt.Printf("%T, %v\n", nums[1:3], nums[1:3])

	// 切片方式赋值方式
	/*
		0. 零值 nil切片
		1. 字面量
		2. make
		3. 切片操作（数组，切片）
	*/
	/*
		切片底层:
			数组地址
			容量
			长度
	*/

	// 长度:len 容量:cap

	nums = []int{}
	fmt.Println(len(nums), cap(nums))
	nums = []int{1, 2, 3}
	fmt.Println(len(nums), cap(nums))
	nums = make([]int, 0)
	fmt.Println(len(nums), cap(nums))
	nums = make([]int, 10)
	fmt.Println(len(nums), cap(nums))
	nums = make([]int, 0, 10)
	fmt.Println(len(nums), cap(nums))

	// numsarray := [5]int{1, 3, 5, 7, 8}
	nums = numsarray[1:3] // length - start, cap - length = 能存储多少个元素
	fmt.Println(len(nums), cap(nums))

	// 元素的访问和修改
	fmt.Println(nums[0]) // 2, 4
	nums[0] = 100        // 索引的范围 0, length - 1
	fmt.Println(nums)
	fmt.Println(nums[1])
	// fmt.Println(nums[2]) // 运行时报错

	// 遍历
	// for 3表达式
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}
	// for range
	for index := range nums {
		fmt.Println(index, nums[index])
	}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	// 如何添加一个元素
	nums = append(nums, 1)
	fmt.Println(nums)

	nums = append(nums, 100, 2, 3)
	fmt.Println(nums)
	// 如何删除一个元素
	// 切片操作
	// 解包
	// 删除索引为0的
	nums = nums[1:len(nums)]
	fmt.Println(nums)
	// 删除索引为length-1的
	nums = nums[0 : len(nums)-1]
	fmt.Println(nums)
	// 删除中间的(i=1)
	// [0:i] + [i+1:len(nums)]
	// prefix := nums[0:1]
	// suffix := nums[2:len(nums)]
	// for _, value := range suffix {
	// 	prefix = append(prefix, value)
	// }
	// fmt.Println(prefix)

	// fmt.Println(append(nums[0:1],
	// 	nums[2:len(nums)][0],
	// 	nums[2:len(nums)][1],
	// 	nums[2:len(nums)][2],
	// 	nums[2:len(nums)][3],
	// 	...,
	// 	nums[2:len(nums)][n]))
	fmt.Println(append(nums[0:1], nums[2:len(nums)]...))

	// start:end start = 0 [:end]
	// start:end end = length [start:]
	// [:]

	// 容量
	// 增长规则： <=1024  n => 2*n
	//  		 > 1024 n => n * (1+~0.25)

	nums = []int{}
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 1)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 2)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 3)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 4)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 5)
	fmt.Println(len(nums), cap(nums))

	// 如果可以预期切片最大长度 可以通过cap指定

	// array[start:end:cap_end]

	numsarray = [5]int{1, 3, 5, 7, 8}
	nums = numsarray[1:3] // length = end - start, cap = arraylength - start
	nums = append(nums, 1000)
	fmt.Println(numsarray) // 1, 3, 5, 7(1000), 8,
	fmt.Println(nums)      // 3, 5, 1000

	nums = numsarray[3:5] // 1000, 8
	nums = append(nums, 3)
	fmt.Println(numsarray) // 1, 3, 5, 1000, 8,
	fmt.Println(nums)      // 1000,8,3

	//[start:end:cap_end] end <=cap_end <= arraylength

	nums = numsarray[1:3:4] //cap = cap_end - start
	fmt.Println(numsarray)
	nums = append(nums, 999)
	fmt.Println(nums)
	fmt.Println(numsarray)
	nums = append(nums, 9999)
	fmt.Println(nums)
	fmt.Println(numsarray)

	// 切片的操作操作
	// [start:end] start<=end<=cap(nums)
	// [start:end:cap_end] start<=end<=cap_end<=cap(nums)
	nums = make([]int, 3, 10)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3

	fmt.Println(nums)

	numsslice := nums[1:2]
	fmt.Println(numsslice)

	fmt.Println(len(numsslice), cap(numsslice))
	numsslice = append(numsslice, 100)
	fmt.Println(nums)

	numsslice = nums[1:2:2]
	fmt.Println(numsslice)
	fmt.Println(len(numsslice), cap(numsslice))
	numsslice = append(numsslice, 999)
	fmt.Println(nums)
	fmt.Println(numsslice)

	// copy函数(dst, src)

	dst := make([]int, 3)
	src := []int{2, 3, 4}

	// len(dst) == len(src)
	fmt.Println(src, dst)
	copy(dst, src)
	fmt.Println(src, dst)

	// len(dst) > len(src)
	src = []int{200, 300}
	copy(dst, src)
	fmt.Println(src, dst) // {200, 300, 4} => 1. 报错, 2.200, 300, 0

	// len(dst) < len(src)
	src = []int{1000, 2000, 3000, 4000}
	copy(dst, src)
	fmt.Println(src, dst) // {1000, 2000, 3000}

	// 删除索引为i的元素

	nums = []int{1, 2, 3, 4, 5}
	// 删除索引为2的
	fmt.Println(nums[2:]) // 3,4,5
	fmt.Println(nums[3:]) //4,5

	copy(nums[2:], nums[3:])

	fmt.Println(nums[:len(nums)-1])

	nums = make([]int, 3, 10)
	nums[0] = 1
	nums[2] = 3
	fmt.Println(nums[:])
}
