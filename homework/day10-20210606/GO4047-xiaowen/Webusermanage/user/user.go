package user

import (
	"fmt"
	"os"
	"strconv"
	// "github.com/princebot/getpass"
	"webusermanage/config"
	"webusermanage/persistence"
)

// var u = map[string]config.UserStructure {
// 	// 1 : {
// 	// 	Id: "1",
// 	// 	UserName: "admin",
// 	// 	Email: "admin@qq.com",
// 	// 	Phone: "123xxxxxxxx",
// 	// 	Password: "123",
// 	// },
// }

// type ServiceUser struct{}



func getId() string{
	id := 0
	u := make(map[string]config.UserStructure)
	if  _, err := os.Stat(config.FilePath); !os.IsNotExist(err){
		p := persistence.Persistence{StroagePath: config.FilePath}
		err := p.Decode(&u)
		if err != nil{
			return err.Error()
		}
		for k,_ := range u{
			n,_ := strconv.Atoi(k)
			if  n > id{
				id = n
			}
		}
		return  strconv.Itoa(id + 1)
	}  

	p := persistence.Persistence{StroagePath: config.FilePath}
	err := p.Decode(&u)
	if err != nil{
		return err.Error()
	}
	return  strconv.Itoa(id + 1) 
}


func Add(user config.UserStructure) (map[string]bool,error){
	u := make(map[string]config.UserStructure)
	p := persistence.Persistence{StroagePath: config.FilePath}
	err := p.Decode(&u)
	if err != nil{
		fmt.Println("add1", err)
		return map[string]bool{"ok":false},err
	}
	n := getId()
	user.Id = n
	u[n] = user
	p = persistence.Persistence{StroagePath: config.FilePath}
	err = p.Encode(u)
	if err != nil{
		fmt.Println("add1", err)
		return map[string]bool{"ok":false},err
	}
	return map[string]bool{"ok":true},nil
}

//3 用户查询
func Query(id string) config.UserStructure{
	u := make(map[string]config.UserStructure)
	// us := make([]byte, 1024)
	p := persistence.Persistence{StroagePath: config.FilePath}
	err := p.Decode(&u)
	if err != nil{
		fmt.Println("unmarsh err:",err)
	}
	for _, v := range u {
		if id == v.Id {
			return v
		}
	}
	return config.UserStructure{}
}

func QueryAll() []config.UserStructure {
	u := make(map[string]config.UserStructure)
	us := make([]config.UserStructure, 0)
	p := persistence.Persistence{StroagePath: config.FilePath}
	err := p.Decode(&u)
	if err != nil{
		fmt.Println("unmarsh err:",err)
	}
	for _, v := range u {
		us = append(us, v)
	}
	// fmt.Println(us)
	return us
}

//4 用户删除
func Delete(id string) map[string]bool{
	u := make(map[string]config.UserStructure)
	p := persistence.Persistence{StroagePath: config.FilePath}
	err := p.Decode(&u)
	if err != nil{
		fmt.Println(err)
	}
	for k, v := range u {
		if id == k {
			delete(u, k)
			fmt.Printf("%s 用户删除成功", v.UserName)
			p := persistence.Persistence{StroagePath: config.FilePath}
			err := p.Encode(u)
			if err != nil{
				fmt.Println("delete", err)
			}
			return map[string]bool{"ok":true}
		}
	}
	return map[string]bool{"ok":false}
}

//5 用户信息修改
func Modif(user *config.UserStructure, id string) map[string]bool {
	u := make(map[string]*config.UserStructure)
	p := persistence.Persistence{StroagePath: config.FilePath}
	err := p.Decode(&u)
	if err != nil{
		return map[string]bool{"ok":false}
	}
	for k, _ := range u {
		if k == id{

			u[k] = user

			
			p := persistence.Persistence{StroagePath: config.FilePath}
			err = p.Encode(u)
			if err != nil{
				fmt.Println("modif1", err)
			}
			return map[string]bool{"ok":true}
		} 
	}
	return map[string]bool{"ok":false}
}

//获取详情
func  GetDtail(id string) (bool, config.UserStructure) {
	u := make(map[string]config.UserStructure)
	// type datas struct{
	// 	Ok bool
	// 	Data config.UserStructure
	// }
	p := persistence.Persistence{StroagePath: config.FilePath}
	err := p.Decode(&u)
	if err != nil{
		fmt.Println("modif1", err)
		// return datas{
		// 	Ok: false,
		// 	Data: config.UserStructure{},
		// }
		return false, config.UserStructure{}
	}
	for k, v := range u {
		if k == id{
			// return datas{
			// 	Ok: true,
			// 	Data: v,
			// }
			return true, v
		} 
	}
	// return datas{
	// 	Ok: false,
	// 	Data: config.UserStructure{},
	// }
	return false, config.UserStructure{}
}