package user

import (
	// "encoding/csv"
	// "encoding/gob"
	"fmt"
	// "io"

	"log"
	// "os"
	// "path/filepath"
	"strconv"
	"strings"

	// "github.com/go-delve/delve/pkg/dwarf/util"
	"github.com/princebot/getpass"
)

//2 用户添加
/* id
username
addr
phone
*/
// var users map[string]map[string]string = map[string]map[string]string{}
// var n string
// var id int = 1
type Useradmin struct{
	Id        string
	UserName  string
	Email     string
	Phone     string
	Password  string
}

func UserAdd(id int, users map[string]map[string]string) map[string]map[string]string {
	var username string
	var phone string
	var addr string
	fmt.Println("请输入用户名:")
	fmt.Scan(&username)
	pwd, err := getpass.Get("请输入密码:")
	if err != nil{
		log.Fatalln("error")
	}
	fmt.Println("请输入地址:")
	fmt.Scan(&addr)
	fmt.Println("请输入联系方式:")
	fmt.Scan(&phone)
	
	users[strconv.Itoa(id)] = map[string]string{
		"username": username,
		"email":    addr,
		"phone": phone,
		"passwd": string(pwd),
	}

	fmt.Println(users)

	id++

	return users
}

//3 用户查询
func Query(users map[string]map[string]string) {
	var q string
	fmt.Println("请输入查询用户:")
	fmt.Scan(&q)
	for k, v := range users {
		if len(q) == 0 {
			fmt.Println(users)
		} else if strings.Contains(v["username"], q) || strings.Contains(v["addr"], q) || strings.Contains(v["phone"], q) {
			fmt.Printf("id:%2s  username:%2s  email:%2s  phone:%2s\n", k, v["username"], v["email"], v["phone"])
		}
	}
}

//4 用户删除
func DelUser(users map[string]map[string]string) {
	var delstr string
	var dete string
	fmt.Println("请输入要删除的用户:")
	fmt.Scan(&delstr)
	for k, v := range users {
		if delstr == v["username"] {
			fmt.Println("是否要删除用户(y/yes/Y/YES):")
			fmt.Scan(&dete)
			if dete == "y" || dete == "yes" || dete == "Y" || dete == "YES" {
				delete(users, k)
			} else {
				return
			}
		}
	}
	fmt.Println(users)
}

//5 用户信息修改
func ModifUser(users map[string]map[string]string) {
	var modstr, orgistr, addrs, phones string
	fmt.Println("请输入要修改的用户名:")
	fmt.Scan(&orgistr)
	for _, v := range users {
		if orgistr == v["username"] {
			fmt.Println("请输入新的用户名:")
			fmt.Scan(&modstr)
			pwd, err := getpass.Get("请输入密码:")
			if err != nil{
				log.Fatalln("error")
			}
			fmt.Println("请输入新地址:")
			fmt.Scan(&modstr)
			fmt.Println("请输入新联系电话:")
			fmt.Scan(&phones)
			v["username"] = modstr
			v["addr"] = addrs
			v["phone"] = phones
			v["passwd"] = string(pwd)
		}
	}
	fmt.Println(users)
}



//序列化
func Serialization(id, username, email, phone, password string){
	var DataType string
	
	fmt.Println("请选择数据持久化类型(gob|csv|json):")
	fmt.Scan(&DataType)
	pre, err := Pre(DataType,id, username, email, phone, password)
	if err !=nil {
		log.Fatalln(err)
	}
	if DataType == "gob"{
		pre.GobPre(id, username, email, phone, password)
	} 
	if DataType == "csv"{
		pre.CsvPre(id, username, email, phone, password)
	}
	if DataType == "json"{
		pre.JsonPre(id, username, email, phone, password)
	}
}
	// switch DataType {
	// case "gob":
	// 	dinfo,err := filepath.Glob("D:\\go\\src\\hello\\usermanage\\user\\*.gob")
	// 	if err != nil{
	// 		log.Fatalln(err)
	// 	}
	// 	if len(dinfo) == 0 {
	// 		gob.Register(&Useradmin{})
	// 		users := []*Useradmin{
	// 			{id, username, email, phone, password},
	// 		}
	// 		file, cerr := os.Create("D:\\go\\src\\hello\\usermanage\\user\\1.gob")
	// 		defer file.Close()
	// 		if cerr != nil{
	// 			log.Fatalln(cerr)
	// 		}
	// 		encoder := gob.NewEncoder(file)
	// 		eerr := encoder.Encode(users)
	// 		if eerr != nil{
	// 			log.Fatalln(eerr)
	// 		}
	// 	} else if len(dinfo) == 4 {
	// 		gob.Register(&Useradmin{})
	// 		users := []*Useradmin{
	// 			{id, username, email, phone, password},
	// 		}
	// 		file, cerr := os.Create("D:\\go\\src\\hello\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".gob")
	// 		defer file.Close()
	// 		reer := os.Remove(dinfo[0])
	// 		if reer != nil{
	// 			log.Fatalln(reer)
	// 		}
	// 		if cerr != nil{
	// 			log.Fatal(cerr)
	// 		}
	// 		encoder := gob.NewEncoder(file)
	// 		eerr := encoder.Encode(users)
	// 		if eerr != nil{
	// 			log.Fatalln(eerr)
	// 		}
	// 	} else{
	// 		gob.Register(&Useradmin{})
	// 		users := []*Useradmin{
	// 			{id, username, email, phone, password},
	// 		}
	// 		file, cerr := os.Create("D:\\go\\src\\hello\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".gob")
	// 		defer file.Close()
	// 		if cerr != nil{
	// 			log.Fatal(cerr)
	// 		}
	// 		encoder := gob.NewEncoder(file)
	// 		eerr := encoder.Encode(users)
	// 		if eerr != nil{
	// 			log.Fatalln(eerr)
	// 		}
	// 	}
	// case "csv":
	// 	dinfo,err := filepath.Glob("D:\\go\\src\\hello\\usermanage\\user\\*.csv")
	// 	if err != nil{
	// 		log.Fatalln(err)
	// 	}
	// 	if len(dinfo) == 0 {
	// 		file, cerr := os.Create("D:\\go\\src\\hello\\usermanage\\user\\1.csv")
	// 		defer file.Close()
	// 		if cerr != nil{
	// 			log.Fatalln(cerr)
	// 		}
	// 		write := csv.NewWriter(file)
	// 		werr := write.Write([]string{id, username, email, phone, password})
	// 		write.Flush()
	// 		if werr != nil{
	// 			log.Fatalln(werr)
	// 		}
	// 	} else if len(dinfo) == 4{
	// 		file, cerr := os.Create("D:\\go\\src\\hello\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".csv")
	// 		defer file.Close()
	// 		reer := os.Remove(dinfo[0])
	// 		if reer != nil{
	// 			log.Fatalln(reer)
	// 		}
	// 		if cerr != nil{
	// 			log.Fatal(cerr)
	// 		}
	// 		write := csv.NewWriter(file)
	// 		werr := write.Write([]string{id, username, email, phone, password})
	// 		write.Flush()
	// 		if werr != nil{
	// 			log.Fatalln(werr)
	// 		}
	// 	} else {
	// 		file, cerr := os.Create("D:\\go\\src\\hello\\usermanage\\user\\"+ strconv.Itoa(len(dinfo)+1)+".csv")
	// 		defer file.Close()
	// 		if cerr != nil{
	// 			log.Fatal(cerr)
	// 		}
	// 		write := csv.NewWriter(file)
	// 		werr := write.Write([]string{id, username, email, phone, password})
	// 		write.Flush()
	// 		if werr != nil{
	// 			log.Fatalln(werr)
	// 		}
	// 	}
	// }


// 加载用户
// func LoadingUser(){
// 	var DataType string
// 	fmt.Println("请输入要加载的数据类型(gob|csv):") 
// 	fmt.Scan(&DataType)
// 	switch DataType{
// 	case "gob":
// 		dinfo,_ := filepath.Glob("D:\\go\\src\\hello\\usermanage\\user\\*.gob")
// 		if  len(dinfo) == 0{
// 			fmt.Println("用户文件还未初始化！！")
// 		} else {
// 			file, err := os.Open(dinfo[len(dinfo)-1])
// 			defer file.Close()
// 			if err != nil{
// 				log.Fatalln(err)
// 			} 
// 			LoadUsers := make([]*Useradmin,0)
// 			gob.Register(&Useradmin{})
// 			dreader := gob.NewDecoder(file)
// 			derr := dreader.Decode(&LoadUsers)
// 			if derr != nil{
// 				fmt.Println(derr)
// 			}
// 		}
// 	case "csv":
// 		dinfo,_ := filepath.Glob("D:\\go\\src\\hello\\usermanage\\user\\*.csv")
// 		if  len(dinfo) == 0{
// 			fmt.Println("用户文件还未初始化！！")
// 		} else {
// 			file, err := os.Open(dinfo[len(dinfo)-1])
// 			defer file.Close()
// 			if err != nil{
// 				log.Fatalln(err)
// 			} 
// 			reader := csv.NewReader(file)
// 			for {
// 				line, derr := reader.Read()
// 				if derr != nil{
// 					if derr == io.EOF{
// 						break
// 					}
// 					fmt.Println(derr)
// 				}
				
// 			}
// 		}
// 	}
// }