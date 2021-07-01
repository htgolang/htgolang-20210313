package models

// type User struct {
// 	ID     int    `json:"id,omitempty"`
// 	Name   string `json:"name,omitempty"`
// 	Passwd string `json:"passwd,omitempty"`
// 	Sex    bool   `json:"sex,omitempty"`
// 	Addr   string `json:"addr,omitempty"`
// 	Tel    string `json:"tel,omitempty"`
// }

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
	Sex    bool   `json:"sex"`
	Addr   string `json:"addr"`
	Tel    string `json:"tel"`
}

func (u *User) SexText() string {
	if u.Sex {
		return "男"
	}
	return "女"
}

func NewUser(id int, name string, passwd string, sex bool, addr, tel string) *User {
	return &User{
		ID:     id,
		Name:   name,
		Sex:    sex,
		Passwd: passwd,
		Addr:   addr,
		Tel:    tel,
	}
}
