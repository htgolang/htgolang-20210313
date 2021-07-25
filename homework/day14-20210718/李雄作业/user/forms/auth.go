package forms

type LoginFrom struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
