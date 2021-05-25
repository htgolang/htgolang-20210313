package login

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/gob"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"github.com/princebot/getpass"
	"path/filepath"
	"log"
	"io"
)

type Useradmin struct{
	Id        string
	UserName  string 
	Email     string
	Phone     string
	Password  string 
}

func Login(u, p string){
	var num = 1
	var user string
	const pwd = `202cb962ac59075b964b07152d234b70`
	fmt.Println("请输入用户名:")
	fmt.Scan(&user)
	pw, err := getpass.Get("请输入密码: ")
	for {
		//得到密码MD5的字符串
		tmp := md5.New()
		tmp.Write(pw)
		tmppw := hex.EncodeToString(tmp.Sum(nil))
		if err != nil {
			fmt.Println(err)
		} 
		//验证
		if user == u && (tmppw == pwd|| string(pw) == p) {
			break
		} else{
			fmt.Println("你输入的用户名或密码有误！！")
			num++
			fmt.Println("请重新输入")
			fmt.Println("请输入用户名:")
			fmt.Scan(&u)
			pw, _ = getpass.Get("请输入密码: ")
		}
		if num == 3{
			os.Exit(1)
		}
	}
}

// 加载用户
func LoadingUser() map[string]map[string]string{
	var DataType string
	fmt.Println("请输入要加载的数据类型(gob|csv|json):") 
	fmt.Scan(&DataType)
	switch DataType{
	case "gob":
		dinfo,_ := filepath.Glob("D:\\go\\src\\hello\\usermanage\\user\\*.gob")
		if  len(dinfo) == 0{
			fmt.Println("用户文件还未初始化！！")
			Login("admin","123")
		} else {
			file, err := os.Open(dinfo[len(dinfo)-1])
			defer file.Close()
			if err != nil{
				log.Fatalln(err)
			} 
			LoadUsers := make([]*Useradmin,0)
			gob.Register(&Useradmin{})
			dreader := gob.NewDecoder(file)
			derr := dreader.Decode(&LoadUsers)
			if derr != nil{
				fmt.Println(derr)
			}
			for _,v := range LoadUsers{
				Login(v.UserName, v.Password)
				
				UserInfo := make(map[string]map[string]string)
				UserInfo[v.Id] = map[string]string{
					"username": v.UserName,
					"email":   v.Email,
					"phone": v.Phone,
					"passwd": v.Password,
				}
				return UserInfo 
			}
		}
	case "csv":
		dinfo,_ := filepath.Glob("D:\\go\\src\\hello\\usermanage\\user\\*.csv")
		if  len(dinfo) == 0{
			fmt.Println("用户文件还未初始化！！")
			Login("admin","123")
		} else {
			file, err := os.Open(dinfo[len(dinfo)-1])
			defer file.Close()
			if err != nil{
				log.Fatalln(err)
			} 
			reader := csv.NewReader(file)
			for {
				line, derr := reader.Read()
				if derr != nil{
					if derr == io.EOF{
						break
					}
					fmt.Println(derr)
				}
				Login(line[1], line[4])
				UserInfo := make(map[string]map[string]string)
				UserInfo[line[0]] = map[string]string{
					"username": line[1],
					"email": line[2],
					"phone": line[3],
					"password": line[4],
				}
				return UserInfo
			}
		}
	case "json":
		dinfo,_ := filepath.Glob("D:\\go\\src\\go-dev\\usermanage\\user\\*.json")
		if  len(dinfo) == 0{
			fmt.Println("用户文件还未初始化！！")
			Login("admin","123")
		} else {
			file, err := os.Open(dinfo[len(dinfo)-1])
			defer file.Close()
			if err != nil{
				log.Fatalln(err)
			} 
			reader := json.NewDecoder(file)

			var u Useradmin
			derr := reader.Decode(&u)
			if derr != nil{
				fmt.Println(derr)
			}
			Login(u.UserName, u.Password)
			UserInfo := make(map[string]map[string]string)
			UserInfo[u.Id] = map[string]string{
				"username": u.UserName,
				"email": u.Email,
				"phone": u.Phone,
				"password": u.Password,
			}
			return UserInfo
		}
	}
	UserInfo := make(map[string]map[string]string)
	return UserInfo
}
