package conf
//package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	GobFilepath     string
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
	filePath := "E:\\002golang\\01NoteAll\\20210424day06\\03homework\\GO4049Arun\\conf\\conf.json"
	filePtr, err := os.Open(filePath)
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