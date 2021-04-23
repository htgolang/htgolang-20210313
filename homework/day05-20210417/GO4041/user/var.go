package user

type user struct {
	id    int
	name  string
	addr  string
	tel   string
	flage bool
}

type Users struct {
	users []user
}

