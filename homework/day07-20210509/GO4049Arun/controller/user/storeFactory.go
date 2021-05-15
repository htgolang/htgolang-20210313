package user

import (
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"mgr/conf"
	"mgr/model/log"
	muser "mgr/model/user"
	"os"
	"path/filepath"
	"strconv"
)


type Persistent interface {
	Write([]*muser.User)
	Read() []*muser.User
}

type jsonPersistent struct {
	absPath string
}
func (j jsonPersistent)Write(user []*muser.User)  {
	file, _ := os.Create(j.absPath)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err := encoder.Encode(user)
	if err == nil{
		fmt.Println("json写入成功!")
	}else {
		log.LogRecord("json文件read失败,请检查原因!")
	}
}

func (j jsonPersistent)Read() []*muser.User {
	file,_ := os.Open(j.absPath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.Token()
	//t,_ := decoder.Token()
	//fmt.Printf("%T:%v\n",t,t)
	var users = make([]*muser.User,0)
	for decoder.More() {
		var _user muser.User
		err := decoder.Decode(&_user)
		if err == nil{
			users = append(users,&_user)
		}else {
			fmt.Println(err)
			log.LogRecord(fmt.Sprintf("json文件read失败,请检查原因:%s",err))
			return nil
		}
	}
	//t,_ = decoder.Token()
	//fmt.Printf("%T:%v\n",t,t)
	return users
}

type csvPersistent struct {
	absPath string
}

func (c csvPersistent)Write(users []*muser.User)  {
	file,_ := os.Create(c.absPath)
	defer file.Close()
	writer := csv.NewWriter(file)
	for i := 0; i < len(users); i++ {
		_tmpInt := int64(users[i].Id)
		writer.Write([]string{strconv.FormatInt(_tmpInt,10),users[i].Name,users[i].Addr,users[i].Tel,users[i].Pwd})
	}
	writer.Flush()
}

func (c csvPersistent)Read() []*muser.User {
	file,_ := os.Open(c.absPath)
	defer file.Close()
	reader := csv.NewReader(file)
	lines,err := reader.ReadAll()
	var users = make([]*muser.User,0)
	if err==nil{
		for i := 0; i <len(lines); i++ {
			var _user muser.User
			_user.Id, _ = strconv.Atoi(lines[i][0])
			_user.Name = lines[i][1]
			_user.Addr= lines[i][2]
			_user.Tel= lines[i][3]
			_user.Pwd= lines[i][4]
			users = append(users,&_user)
		}
	}else {
		log.LogRecord(fmt.Sprintf("Csv文件read失败,请检查原因!%v",err))
	}
	return users
}


type gobPersistent struct {
	absPath string
}

func (g gobPersistent)Write(users []*muser.User)  {
	file, _ := os.Create(g.absPath)
	//(1)写入
	encoder := gob.NewEncoder(file)
	err := encoder.Encode(users)
	file.Close()
	if err==nil{
		fmt.Println("Gob写入成功")
	}else {
		log.LogRecord(fmt.Sprintf("Gob文件写入失败,请检查原因!%v",err))
	}
}

func (g gobPersistent)Read() []*muser.User {
	fileDecode, _ := os.Open(g.absPath)
	defer fileDecode.Close()
	decoder := gob.NewDecoder(fileDecode)
	users2 := make([]*muser.User, 0)
	//Decode存储到users2中
	err := decoder.Decode(&users2)
	if err == nil{
		return users2
	}else{
		log.LogRecord("Parse err!")
		return nil
	}
	return nil
}

func Create(s string) Persistent {
	switch s {
	case "json":
		absPath , _:= filepath.Abs(conf.Info.JsonFilepath)
		return &jsonPersistent{absPath}
	case "csv":
		absPath , _:= filepath.Abs(conf.Info.CsvFilepath)
		return &csvPersistent{absPath}
	case "gob":
		gob.Register(&muser.User{})
		absPath , _:= filepath.Abs(conf.Info.GobFilepath)
		return &gobPersistent{absPath}
	}
	return nil
}
