package config

var FilePath string = "D:\\go\\src\\go-dev\\Webusermanage\\stroage\\user.json"

type UserStructure struct{
	Id        string `json:"id"`
	UserName  string  `json:"username"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Password  string  `json:"password"`
}

// U := map[int]&UserStructure{
// 	1:{

// 	}
// }
type U *UserStructure


