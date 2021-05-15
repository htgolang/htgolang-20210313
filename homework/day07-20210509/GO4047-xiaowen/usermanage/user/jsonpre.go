package user

type JsonData struct{
	Presistnece
}

func NewJson(id, username, email, phone, password string) PreType {
	return &JsonData{
		Presistnece: Presistnece{
			Id:  id,
			Username: username,
			Email: email,
			Phone: phone,
			Password: password,
		},
	} 
}