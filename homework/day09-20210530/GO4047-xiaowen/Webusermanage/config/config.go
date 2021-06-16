package config

var FilePath string = "D:\\go\\src\\go-dev\\Webusermanage\\stroage\\user.json"

type UserStructure struct{
	Id        string
	UserName  string
	Email     string
	Phone     string
	Password  string
}

// U := map[int]&UserStructure{
// 	1:{

// 	}
// }
type U *UserStructure


