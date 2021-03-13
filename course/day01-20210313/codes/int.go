package main

import "fmt"

func main() {
	var (
		i1 int = 17         //10进制
		i2 int = 021        //8进制
		i3 int = 0x11       // 16进制
		i4 int = 0b00010001 // 2进制
	)
	fmt.Printf("i1:%T, i1:%v\n", i1, i1)
	fmt.Printf("i2:%T, i2:%v\n", i2, i2)
	fmt.Printf("i3:%T, i3:%v\n", i3, i3)
	fmt.Printf("i4:%T, i4:%v\n", i4, i4)
	// 0b00001111
	// 1 * 2 ^ 3 + 1 * 2 ^2+1 * 2 ^ 1 + 1 * 2^0
	// 8 + 4 + 2 + 1 => 15

	var a byte = 'A' // a = 65
	fmt.Printf("%T, %v\n", a, a)

	var u rune = '我' //25105
	fmt.Printf("%T, %v\n", u, u)

	var l int64 = 1
	fmt.Printf("%T, %v\n", l, l)

	// 左操作数和右操作数类型必须一致
	fmt.Println(i1+i2, i1-i2, i1*i2, i1/i2, i1%i2)
	i1++ // i1 = i1 + 1
	i2-- // i2 = i2 - 1
	fmt.Println(i1, i2)
	// 除法
	fmt.Println(5 / 2) // => int
	// x := 0
	// fmt.Println(5 / x)
	fmt.Println(1 < -1)

	// 输出位二进制
	fmt.Printf("%b\n", i3)
	fmt.Printf("%b\n", 0b0010|0b1010)  //1010
	fmt.Printf("%b\n", 0b0110&^0b1010) //0100

	i3 += 3 // i3 = i3 + 3
	fmt.Println(i3)

	// l => int64
	// i3 => int
	// 强制类型转换 l => int? i3 => int64
	// 小范围-> 大范围 肯定无问题
	// 大范围->小范围 可能被截断
	fmt.Println(int(l) + i3)
	fmt.Println(l + int64(i3))

	var i6 int16 = 255
	fmt.Println(int8(i6)) // -128-127

	// fmt
	fmt.Printf("%d, %#b, %#o, %#x\n", i4, i4, i4, i4)
	fmt.Printf("%c %c\n", a, u)
	fmt.Printf("%U %U\n", a, u)
	fmt.Printf("%q %q\n", a, u)

	// %d 修饰符
	fmt.Printf("%+d, %+d\n", 1, -1)
	fmt.Printf("%5d, %5d\n", 1, -1)
	fmt.Printf("%5d, %5d\n", 655356, -655356)
	fmt.Printf("%05d, %05d\n", 1, -1)
	fmt.Printf("%-5d, %-5d\n", 1, -1)

	fmt.Println(1_1) // 11
	// 1_1 数字之间可以加_

	fmt.Println(0123, 123)
}
