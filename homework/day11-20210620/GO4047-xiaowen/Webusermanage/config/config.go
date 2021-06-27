package config

var DSN string = "root:123.com@tcp(192.168.221.130:3306)/mycmdb?charset=utf8mb4&parseTime=true"
type UserStructure struct{
	Id        string `json:"id"`
	UserName  string  `json:"username"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Password  string  `json:"password"`
	Role_id   int
}
var Addr = ":9999"




