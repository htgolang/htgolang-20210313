package models

type User struct {
	ID   int
	Name string
	Sex  bool
}

func (u *User) SexText() string {
	if u.Sex {
		return "男"
	}
	return "女"
}
