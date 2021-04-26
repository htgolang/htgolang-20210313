package model

type Persistence struct {
	Types string
}

func (p *Persistence) Decode(d func() (map[int]User, error)) (map[int]User, error) {
	return d()
}
func (p *Persistence) Encode(e func(map[int]User) error, user map[int]User) error {
	return e(user)
}
