package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// 属性名称都大写
// 序列化和反序列化需要访问属性
type User struct {
	ID   int
	Name string
	Addr string
}

func main() {
	// 分go内置的基本类型和复合数据类型
	gob.Register(&User{})

	users := []*User{
		{1, "kk", "西安"},
		{2, "alinx", "深圳"},
	}

	file, _ := os.Create("data.gob")

	// 编码
	encoder := gob.NewEncoder(file)

	err := encoder.Encode(users)
	fmt.Println(err)
	file.Close()

	file, _ = os.Open("data.gob")
	defer file.Close()

	// 解码
	decoder := gob.NewDecoder(file)

	users2 := make([]*User, 0)

	err = decoder.Decode(&users2)

	fmt.Println(err, users2)
	for _, u := range users2 {
		fmt.Printf("%#v\n", u)
	}
}
