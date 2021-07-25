package forms

type AddUserForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Addr     string `form:"addr"`
	Sex      bool   `form:"sex"`
}

type ChangeUserForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Addr     string `form:"addr"`
	Sex      bool   `form:"sex"`
}
