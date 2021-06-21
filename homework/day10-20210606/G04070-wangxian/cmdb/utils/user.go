package utils

import (
	"cmdb/modules"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func LoadUserInfo(path string) ([]modules.User, error) {
	var users []modules.User

	file, err := os.Open(path)
	if err != nil {
		log.Printf("open user info file %s faild. %s", path, err)
		return nil, err
	}

	decoder := json.NewDecoder(file)
	decoder.Decode(&users)
	return users, nil
}

func SaveUserInfo(path string, users []modules.User) error {
	os.Rename(path, path+".bak")
	file, err := os.Create(path)
	if err != nil {
		log.Printf("create new user info file %s faild. %s", path, err)
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(users)
	return nil
}

func Md5Encrypt(b []byte) string {
	hasher := md5.New()
	hasher.Write(b)
	hasher.Write([]byte("asxeft"))
	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	return hash
}

//检查用户名是否冲突
func CheckUserName(name string, users []modules.User) bool {
	for _, v := range users {
		if name == v.Name {
			return true
		}
	}
	return false
}

//检查号码是否冲突
func CheckTel(tel string, users []modules.User) bool {
	for _, v := range users {
		if tel == v.Tel {
			return true
		}
	}
	return false
}

func GenerateID(users []modules.User) int {
	id := 0
	for _, v := range users {
		i := v.Id
		if i > id {
			id = i
		}
	}
	return id + 1
}
