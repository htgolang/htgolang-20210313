package main

import (
	"encoding/json"
	"fmt"
)

/* json: 格式化的字符串
数值：number(int/float) =>  1.1 10 -10
字符串：string => ""
布尔：bool => true/false
数组：[]Type/[Length]Type => [1, 2, 3, "", ""]
字典(Object)：map[KType]Value => {"key1":value, "key2":value }

Go数据类型 <=> Gob二进制字符串
Go数据类型 <=> Json字符串
Go数据类型 <=> xml字符串
Go数据类型 <=> Protobuffer

持久化/网络传输
=> 序列化 => Go数据类型 => 字符串(文本/二进制)
=> 反序列化 => 字符串(文本/二进制) = >Go数据类型

编码规则 => 库(包)
*/
func main() {
	// json
	var (
		number float64        = 1.1
		name   string         = "kk"
		isBody bool           = true
		users  []string       = []string{"arue", "方平", "刘冠威"}
		scores map[string]int = map[string]int{"arune": 90, "方平": 80, "蓝海滨": 80}
	)

	b, err := json.Marshal(number)
	fmt.Printf("%#v, %T, %#v\n", err, b, string(b))

	b, err = json.Marshal(name)
	fmt.Printf("%#v, %T, %#v\n", err, b, string(b))

	b, err = json.Marshal(isBody)
	fmt.Printf("%#v, %T, %#v\n", err, b, string(b))

	b, err = json.Marshal(users)
	fmt.Printf("%#v, %T, %#v\n", err, b, string(b))

	b, err = json.Marshal(scores)
	fmt.Printf("%#v, %T, %#v\n", err, b, string(b))

	jsonB := `
	{
		"name" : "kk",
		"age" : 30,
		"isBody" : true,
		"scores" : [1, 2, 3, 4, 5]
	}
	`
	var rs map[string]interface{}

	err = json.Unmarshal([]byte(jsonB), &rs)
	fmt.Printf("%#v, %#v\n", err, rs)

}
