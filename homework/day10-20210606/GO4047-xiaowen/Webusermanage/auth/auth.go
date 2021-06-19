package auth

import (
	"webusermanage/config"
	"webusermanage/persistence"
)

func Auth(u, pwd string) (string, bool){
	user := make(map[string]config.UserStructure)
	p := persistence.Persistence{StroagePath: config.FilePath}
	err := p.Decode(&user)
	if err != nil{
		return "", false
	}
	for _, v := range user{
		if u == v.UserName && pwd == v.Password{
			return v.Id, true
		}
	}
	return "",false
}