package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)
type Useradmin struct{
	Id        string
	UserName  string 
	Email     string
	Phone     string
	Password  string 
}
func main(){
	dinfo,_ := filepath.Glob("D:\\go\\src\\go-dev\\usermanage\\user\\*.json")
	fmt.Println(dinfo)
		if  len(dinfo) == 0{
			fmt.Println("用户文件还未初始化！！")
		} else {
			file, err := os.Open(dinfo[len(dinfo)-1])
			defer file.Close()
			if err != nil{
				log.Fatalln(err)
			} 
			b, err := ioutil.ReadAll(file)
			if err != nil{
				fmt.Println(err)
			}

			
			reader := json.NewDecoder(file)

			var u Useradmin
			d := json.Unmarshal(b, &u)
			derr := reader.Decode(&u)
			if derr != nil{
				fmt.Println(derr)
			}
			fmt.Printf("%#v, %#v", u, d)
			UserInfo := make(map[string]map[string]string)
			UserInfo[u.Id] = map[string]string{
				"username": u.UserName,
				"email": u.Email,
				"phone": u.Phone,
				"password": u.Password,
			}
		}	
}