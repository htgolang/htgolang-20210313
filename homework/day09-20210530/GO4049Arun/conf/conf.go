package conf
//package main
import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type config struct {
	GobFilepath     string
	JsonFilepath     string
	CsvFilepath     string
	LogFilepath 	string
	PwdLen			int
}

func parseJson(filePtr *os.File,info config) config {
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	decoder.Decode(&info)
	//err := decoder.Decode(&info)
	//if err != nil {
	//	fmt.Println("解码失败", err.Error())
	//} else {
	//	fmt.Println("解码成功")
	//}
	return info
}
var Info config

func init() {
	filePath := "conf/conf.json"
	absPath , _:= filepath.Abs(filePath)
	fmt.Println(absPath)
	filePtr, err := os.Open(absPath)
	defer filePtr.Close()
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	Info  = parseJson(filePtr,Info)
	//fmt.Println(Info.GobFilepath, Info.LogFilepath)
	//fmt.Printf("%T",Info.PwdLen)
	//model/store/user.gob    model/log
}