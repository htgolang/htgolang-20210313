package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var srcfile,dstfile =os.Args[1],os.Args[2]

func main(){

	start := time.Now()

	conts :=myfile()
	file, err := os.OpenFile(dstfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	//写入字节切片数据
	file.WriteString(conts) //直接写入字符串数据

	//计算copy文件的时间
	fmt.Printf("复制用时%v\n", time.Now().Sub(start))
}
func myfile() string{

	content, err := ioutil.ReadFile(srcfile)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return ""
	}

	return string(content)

}

//1K     复制用时567.4µs
//1M     复制用时2.0256ms
//10M    复制用时19.7831ms