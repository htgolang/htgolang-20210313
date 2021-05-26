package admin

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type JsonClient struct {
}

func NewJsonClient() *JsonClient {
	return &JsonClient{}
}

// 在Json文件中读取用户文件
func (c *JsonClient) ReadUsersData() error {
	// 判断用户文件是否存在，如果初始用户文件不存在，则写入一个初始用户进文件内
	c.CheckUserDataFileExist(UserDataJsonFile)

	// 读取用户数据
	file, err := os.Open(UserDataJsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var u Users
		err := json.Unmarshal([]byte(line), &u)
		if err != nil {
			return err
		}
		UserList = append(UserList, &u)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Cannot scanner text file: %s, err: [%v]", UserDataJsonFile, err)
		return err
	}
	return nil
}

// 用户信息写入到Json文件中
func (c *JsonClient) WritesUsersData() error {
	file, err := os.Create(UserDataJsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, v := range UserList {
		u, err := json.Marshal(v)
		if err != nil {
			return err
		}
		_, err = file.Write(u)
		if err != nil {
			panic(err)
		}
	}
	return err
}

// 检查初始用户文件是否存在
func (c *JsonClient) CheckUserDataFileExist(f string) error {
	// 判断用户文件是否存在，如果初始用户文件不存在，则写入一个初始用户进文件内
	_, err := os.Stat(f)
	if err != nil {
		user := &Users{
			Id:     InitialUser[0],
			Name:   InitialUser[1],
			Passwd: InitialUser[2],
			Addr:   InitialUser[3],
			Tel:    InitialUser[4],
		}
		// 序列化初始用户
		u, err := json.Marshal(user)
		if err != nil {
			return err
		}
		file, err := os.Create(f)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.Write(u)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

// 备份用户文件
func (c *JsonClient) CopyUserDataFile() error {
	// 源数据文件
	source, err := os.Open(UserDataJsonFile)
	if err != nil {
		return err
	}
	defer source.Close()

	// 备份数据文件，已时间格式结尾
	dstFile := UserDataDirJson +
		UserDataFileStr +
		"-" + time.Now().Format("2006-01-02_150405") +
		UserDataJsonFileFormat
	destination, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	// 复制源文件到备份文件
	_, err = io.Copy(destination, source)
	return err
}

// 持久化最近N次修改的文件，超出的部分删除最旧的
func (c *JsonClient) PersistenceOfLastNChanges() error {
	file, err := filepath.Glob(UserDataDirJson + UserDataFileStr + "*")
	if err != nil {
		return err
	}
	if len(file) > KeepUserDataFileNum {
		err = os.Remove(file[0])
		if err != nil {
			return err
		}
	}
	return err
}
