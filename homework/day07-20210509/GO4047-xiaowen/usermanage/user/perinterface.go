package user

type PreType interface{
	CsvPre(id, username, email, phone, password string)
	GobPre(id, username, email, phone, password string)
	JsonPre(id, username, email, phone, password string)
}

// type preDtat struct{}

// func (t preDtat) Gob(){}
// func (t preDtat) Csv(){}
// func (t preDtat) Json(){}

