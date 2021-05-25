package user


type CsvData struct{
	Presistnece
}

func NewCsv(id, username, email, phone, password string) PreType {
	return &CsvData{
		Presistnece: Presistnece{
			Id:  id,
			Username: username,
			Email: email,
			Phone: phone,
			Password: password,
		},
	}
}