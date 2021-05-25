//1. copyfile 测试bufferSize大小对程序执行速度影响

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)


//文件大小：200M
func main(){
	start := time.Now().Unix()
	file,err := os.Open("F:\\BaiduNetdiskDownload\\CRUSH运行图调制入门(01).exe")
	defer file.Close()
	if err != nil{
		fmt.Println(err)
	}
	// data := make([]byte, 1024*1024)  用时:  1
	data := make([]byte, 1024*1024*10)  // 用时:  0 (小于1秒，接近于0)
	dest,derr := os.Create("F:\\BaiduNetdiskDownload\\CRUSH.exe")
	if derr != nil{
		fmt.Println(derr)
	}
	for {
		n,nerr := file.Read(data)
		if nerr != nil{
			if nerr == io.EOF{
				break
			}
			fmt.Println(nerr)
		}
		dest.Write(data[:n])
	}
	fmt.Println("用时: ", time.Now().Unix() - start)
}