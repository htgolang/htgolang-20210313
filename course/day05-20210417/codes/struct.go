package main

import (
	"fmt"
	"time"
)

// 定义结构体
/*
type StructName struct {
	Field1 FieldType1
	Field2 FieldType2
	...
	Fieldn FieldTypen
}


user:
	id => int
	name => string
	addr => string
	tel => string
	birthday => string/uinxtime/time.Time => 首选time
	created_at => time.Time/string/unixtime => 首选time
	updated_at => time.Time
	flag => bool // true 删除, false未删除
	status => int // 1删除, 0未删除
	deleted_at => *time.Time // 已删除:非 nil, 未删除 nil // time.Time 已删除:非零值(删除时间) 未删除: 时间零值


users = [] //真正从切片中移除掉
物理删除: 真删除
逻辑删除: 假删除
	通过标识，标识数据被删除，原有数据还存在
users => 逻辑删除 map[flag] => "0" / "1"

*/

// 定义结构体User => 类 => 封装
type User struct {
	id        int
	name      string
	addr      string
	tel       string
	birthday  time.Time
	createdAt time.Time
	updatedAt time.Time
	flag      bool
	status    int
	deletedAt *time.Time
}

// 类 构造器 => 创建对应类的实例(对象/变量) // Go中无构造器的说法
// 函数 N(n)ew结构体名称 构建结构体的（值/指针）对象

func NewUser(id int, name, addr, tel string, birthday time.Time) User {
	return User{
		id:        id,
		name:      name,
		addr:      addr,
		tel:       tel,
		birthday:  birthday,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		flag:      false,
		status:    0,
		deletedAt: nil,
	}
}

func NewUserPointer(id int, name, addr, tel string, birthday time.Time) *User {
	return &User{
		id:        id,
		name:      name,
		addr:      addr,
		tel:       tel,
		birthday:  birthday,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		flag:      false,
		status:    0,
		deletedAt: nil,
	}
}

type User2 struct {
	id              int
	name, addr, tel string
}

func main() {
	// 指针
	var t *time.Time
	fmt.Println(t) // 零值 nil

	// 定义User类型的变量
	var u User // 由所有属性的零值组成的一个User类型的变量值

	fmt.Printf("%T\n", u)
	// 零值 => 由所有属性的零值组成的变量值
	fmt.Printf("%v\n", u) // 非nil
	fmt.Printf("%#v\n", u)

	// 赋值 字面量
	// 设置值 按照 结构体中属性定义的顺序设置
	// 需要给所有属性都进行赋值
	// 按照属性顺序定义的字面量
	u = User{1, "kk", "西安", "152xxxxxxxxx", time.Now(), time.Now(), time.Now(), false, 0, nil}

	fmt.Printf("%#v\n", u)
	// 按照属性名称定义的字面量
	// 未指定的属性名称 对应类型的0值进行初始化
	u = User{id: 2, name: "陆建峰"}

	fmt.Printf("%#v\n", u)

	// 属性进行访问和修改

	fmt.Println(u.id)
	fmt.Println(u.name)
	u.addr = "北京"

	fmt.Printf("%#v\n", u)

	// 值类型
	u2 := u
	u2.tel = "152xxxxx"
	fmt.Printf("%#v\n", u2)
	fmt.Printf("%#v\n", u) // 不影响

	// 指针类型的对象
	var u3 *User

	fmt.Printf("%T, %v\n", u3, u3)
	// u3 =
	// 指针赋值 => 取引用 &
	// new 申请空间并对元素使用0值进行初始化，取地址，赋值
	u3 = &u2
	fmt.Printf("%#v\n", u3)

	var u4 *User = new(User)

	fmt.Printf("%#v\n", u4)
	// 字面量取引用 User{}
	var u5 *User = &User{}

	fmt.Printf("%#v\n", u5)

	u5 = &User{id: 10}

	fmt.Printf("%#v\n", u5)

	// u = User{}

	// fmt.Printf("%#v\n", u)

	u6 := NewUser(10, "黄明", "杭州", "", time.Now())
	u7 := NewUserPointer(11, "张松", "", "", time.Now())

	fmt.Printf("%T, %#v\n", u6, u6)
	fmt.Printf("%T, %#v\n", u7, u7)

	var user2 User2 = User2{
		1,
		"kk", "xxx", "yyyy",
	} //

	fmt.Printf("%T, %#v\n", user2, user2)
}
