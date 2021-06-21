package modules

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Sex    bool   `json:"sex"`
	Addr   string `json:"addr"`
	Tel    string `json:"tel"`
	Passwd string `json:"passwd"`
}

func (u User) SetText() string {
	if u.Sex {
		return "男"
	}
	return "女"
}
