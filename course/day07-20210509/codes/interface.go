package main

import "fmt"

type User struct {
	Id   int
	Name string
}

// 定义接口Persistent
// 行为: Save保存User切片, Load加载User切片
type Persistent interface {
	Save([]User, string) error
	Load(string) ([]User, error)
}

type GobPersistent struct{}

// 定义了Persistent 所有的接口，就叫GobPersistent实现了Persistent接口

func (p GobPersistent) Save(users []User, path string) error {
	fmt.Println("gob persistent save")
	return nil
}

func (p GobPersistent) Load(path string) ([]User, error) {
	fmt.Println("gob persistent load")
	return nil, nil
}

type CsvPersistent struct{}

func (p CsvPersistent) Save(users []User, path string) error {
	fmt.Println("csv persistent save")
	return nil
}

func (p CsvPersistent) Load(path string) ([]User, error) {
	fmt.Println("csv persistent load")
	return nil, nil
}

func call(persistent Persistent) {
	fmt.Println("call:")
	persistent.Save(nil, "")
	persistent.Load("")
}

type JsonPersistent struct{}

func (p *JsonPersistent) Save(users []User, path string) error {
	fmt.Println("json persistent save")
	return nil
}

func (p *JsonPersistent) Load(path string) ([]User, error) {
	fmt.Println("json persistent load")
	return nil, nil
}

type ProtobufPersistent struct{}

func (p *ProtobufPersistent) Save(users []User, path string) error {
	fmt.Println("protobuf persistent save")
	return nil
}

func (p ProtobufPersistent) Load(path string) ([]User, error) {
	fmt.Println("protobuf persistent load")
	return nil, nil
}

func main() {
	var persistent Persistent
	fmt.Printf("%T, %#v\n", persistent, persistent)
	// 赋值
	// 接口 不能直接通过接口类型初始化对象
	// 需要使用实现行为（实现接口的所有方法，某个对象的类型定义了接口中所有的方法）的对象

	persistent = CsvPersistent{} // 如果有接口后只需要修改实现

	fmt.Printf("%T, %#v\n", persistent, persistent)
	persistent.Save(nil, "") // => GobPersistent->Save
	persistent.Load("")      // => GobPersistent->Load
	call(persistent)

	// 多态 => 某个对象赋值为不同的对象 体现出不同的行为
	fmt.Println("多态")
	fmt.Println("Csv:")
	call(CsvPersistent{})
	fmt.Println("Gob:")
	call(GobPersistent{})

	// 定义类型 与接口 是否实现 无语法上直接关联
	// 鸭子类型 你怎么判断一个动物是否是鸭子
	// => 看鸭子的行为: 游泳，嘎嘎叫
	// 可以游泳的并且会嘎嘎叫的都是鸭子

	// 指针对象
	persistent = new(CsvPersistent)
	fmt.Printf("%T, %#v\n", persistent, persistent)
	persistent.Load("")
	persistent.Save(nil, "")

	call(&CsvPersistent{})
	call(&GobPersistent{})
	// 疑问 为什么指针对象也有Save和Load方法
	// GO自动为值接收者方法生成指针接收者方法

	// 是语法糖吗？ 不是
	// 值接收者 指针.值接收者方法() => 语法糖 解引用

	// 类型指针接收者方法 如何赋值

	// JsonPersistent{} = 不可以 => 未实现值接收者方法
	// new(JsonPersistent) = 可以
	fmt.Println("指针接收者:")
	persistent = new(JsonPersistent)
	persistent.Load("")
	persistent.Save(nil, "")

	// persistent = JsonPersistent{}
	// persistent.Load("")
	// persistent.Save(nil, "")

	// 又有指针，又有值
	// ProtobufPersistent{} => 未实现save值接收者方法
	// new(ProtobufPersistent)
	// persistent = ProtobufPersistent{}
	// persistent.Load("")
	// persistent.Save(nil, "")

	persistent = &ProtobufPersistent{}
	persistent.Load("")
	persistent.Save(nil, "")

}
