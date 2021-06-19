package modules

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Addr   string `json:"addr"`
	Tel    string `json:"tel"`
	Passwd string `json:"-"`
}
