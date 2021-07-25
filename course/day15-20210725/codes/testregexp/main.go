package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 正则
	// regexp

	// 判断是否匹配
	//		matched, err := regexp.MatchString
	//		regexp.Match
	//      regexp.MatchReader
	//	[]byte, Reader, String
	// 查找字符串 => 找到匹配的字符串
	// 替换字符串
	// 分隔字符串

	// 正则表达式的书写
	// 开始 ^
	// 结尾 $
	// 集
	// 		[a-z]
	// 		[a-zA-Z]
	// 		[a-zA-Z0-9]
	// 		\d 0-9
	// 		\D \d取反 [^0-9]
	// 		\w 0-9a-zA-Z_
	// 		\W \w取反 [^0-9a-zA-Z_]
	// 		. 任意字符
	// 数量限定符
	// 		0或1个     ?
	// 		至少1个	  +
	// 		任意多个	 *
	// 		最少n个	{n,}
	// 		最多m个	{,m}
	// 		n个到m个	{n,m}
	// 分组
	// 			()
	// 特殊字符转义
	// 		\
	// [\uxxxx-\uxxxx]
	// 字符串 int  1-10长度的只能包含数字 0-9999999999

	text := "100"
	// "\\d{1,10}"
	// "[0-9]{1,10}"
	// `\d{1,10}`

	// 验证长度 字符串
	matched, err := regexp.MatchString("^[0-9]{1,10}$", text)
	fmt.Println(matched, err)
	// 修改用户名
	// 包含大小写英文字母数字_ 长度3-32
	// `\w{3,32}`
	// 验证邮箱 大小写英文字母数字@大小写英文字母数字.大小写英文字母数字
	// \w*@\w*\\.\w*

	// .com
	// .net
	// .cn
	// 或者 |
	matched, err = regexp.MatchString(`^\w{1,32}@\w*[.]((com)|(net)|(cn))$`, "test@test.com")
	fmt.Println(matched, err)

	// 查找字符串中的邮箱
	text = "我的邮箱是kk@test.com, 我的另外一个邮箱是kk@xxx.cn, 我的qq邮箱:kk@test.qq"
	reg := regexp.MustCompile(`\w{1,32}@\w*[.]((com)|(net)|(cn))`)
	fmt.Println(reg.FindString(text))
	fmt.Println(reg.FindAllString(text, -1))

	// 替换
	fmt.Println(reg.ReplaceAllString(text, "*@*.*"))

	// 分隔
	// ip;,空白符ip
	emptyReg := regexp.MustCompile(`[;,\s]`)
	fmt.Printf("%#v\n", emptyReg.Split("1.1.1.1;2.2.2.2,3.3.3.3 4.4.4.4\t5.5.5.5", -1))
}
