package admin

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type CsvClient struct {
}

func NewCsvClient() *CsvClient {
	return &CsvClient{}
}

// 在csv文件中读取用户文件
func (c *CsvClient) ReadUsersData() {
	// 判断用户文件是否存在，如果初始用户文件不存在，则写入一个初始用户进文件内
	c.CheckUserDataFileExist()

	// 读取用户数据
	file, _ := os.Open(UserDataCsvFile)
	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		//添加用户数据，迁移至user.go文件中进行所有与数据库有关的操作
		var (
			// 根据输入值插入数据
			u = &Users{
				Id:     line[0],
				Name:   line[1],
				Passwd: line[2],
				Addr:   line[3],
				Tel:    line[4],
			}
		)
		UserList = append(UserList, u)
	}
}

// 用户信息写入到CSV文件中
func (c *CsvClient) WritesUsersData() error {
	//User status
	userStatus, err := os.Create(UserDataCsvFile)
	if err != nil {
		return err
	}
	writer1 := csv.NewWriter(userStatus)
	for _, v := range UserList {
		err = writer1.Write(
			[]string{
				v.Id,
				v.Name,
				v.Passwd,
				v.Addr,
				v.Tel,
			})
		if err != nil {
			return err
		}
		writer1.Flush()
	}
	return err
}

// 检查初始用户文件是否存在
func (c *CsvClient) CheckUserDataFileExist() error {
	// 判断用户文件是否存在，如果初始用户文件不存在，则写入一个初始用户进文件内
	_, err := os.Stat(UserDataCsvFile)
	if err != nil {
		userStatus, err := os.Create(UserDataCsvFile)
		if err != nil {
			return err
		}
		writer1 := csv.NewWriter(userStatus)
		_ = writer1.Write(InitialUser)
		writer1.Flush()
	}
	return err
}

// 备份用户文件
func (c *CsvClient) CopyUserDataFile() error {
	// 源数据文件
	source, err := os.Open(UserDataCsvFile)
	if err != nil {
		return err
	}
	defer source.Close()

	// 备份数据文件，已时间格式结尾
	dstFile := UserDataDirCsv +
		UserDataFileStr +
		"-" + time.Now().Format("2006-01-02_150405") +
		UserDataCsvFileFormat
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
func (c *CsvClient) PersistenceOfLastNChanges() error {
	file, err := filepath.Glob(UserDataDirCsv + UserDataFileStr + "*")
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
