package models

// 包内可见(数据封装 数据对外隐藏)
type userv1 struct {
	id   int
	Name string
}

// 在包外 需要使用userv1类型的变量
func NewUserv1(id int, name string) userv1 {
	return userv1{id, name}
}

// 在包外想要修改ID的值
func SetUserv1Id(u userv1, id int) {
	u.id = id
}

func SetUserv1IdP(u *userv1, id int) {
	u.id = id
}

// 包外可见
type Userv2 struct {
	id   int    // 首字母小写 包内可见
	Name string // 首字母大写 包外可见
}

type Userv3 struct {
	User userv1
	Desc string
}

/*
包外
Userv3 => 可以访问
Userv3.User => 可以访问
Userv3.User.id => 不可以访问
Userv3.User.Name => 可以访问
Userv3.Desc => 可以访问
*/

type Userv4 struct {
	userv1
	Desc string
}

/*
Userv4 => 可以访问
Userv4.userv1 => 不能访问 属性名小写
Userv4.userv1.id => 不能访问
Userv4.id => 不行
Userv4.userv1.Name => 不能访问
Userv4.Name => 可以访问
Userv4.Desc => 可以访问
*/
