package user

type GobData struct{
	Presistnece
}

func NewGob(id, username, email, phone, password string) PreType {
	return &GobData{
		Presistnece: Presistnece{
			Id:  id,
			Username: username,
			Email: email,
			Phone: phone,
			Password: password,
		},
	}
}