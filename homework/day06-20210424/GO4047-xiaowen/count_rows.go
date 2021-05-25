// 2. 代码行数统计
//     输入: 目录/文件 文件后缀
//     输出: Go程序的代码行数(空行, }, #)

//     思路：递归，统计\n数量，带缓冲区IO

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)


func getAllFile(args string) ([]string, error){

	// args can pass three types of arguments: dricetory| file| file suffix
	// For example: 
	//    	count_rows.go test/
	// 		count_rows.go hh.txt
	//		count_rows.go "./*.txt"  ##在计算go文件时需要编译后才能执行成功

	var Allfiles []string

	// 判断传递的是否后缀类型
	suf := []string{"*.go","*.txt","*.log","*.py","*.awk","*.sh"}
	for _,v := range suf {
		if strings.Contains(args, v) {
			ma, merr := filepath.Glob(args)
			if merr != nil{
				log.Println("1",merr)
				return nil, merr
			} else {
				return ma,merr
			}
		}
	}

	// 判断传递的是否为文件
	_, ferr := os.Stat(args)
	if ferr != nil{
		log.Println("ferr",ferr)
	} else {
		Allfiles = append(Allfiles, args)
		return Allfiles,nil
	}

	// 判断传递的是否为目录
	fileinfo, err := ioutil.ReadDir(args) 
	if err != nil{
		log.Println("2",err)
		return nil, err
	}
	for  _,v := range fileinfo {
		if v.IsDir(){
			newf, nerr := getAllFile(args+"\\"+v.Name())
			
			if nerr != nil {
				log.Println("3",nerr)
				return nil,nerr
			} 
			Allfiles = append(Allfiles, newf...)
	
		} else {
			Allfiles = append(Allfiles, args+"\\"+v.Name())
		}

	}
	return Allfiles,nil
}



func main(){

	num := 0
	if len(os.Args) != 2{
		fmt.Println(len(os.Args))
		log.Println("Usage: count_rows.go dricetory|file|suffix")
	}
	files, err := getAllFile(os.Args[1])
	if err != nil{
		log.Println("4",err)
	}
	
	for _,v := range files{
		ofile, oerr := os.Open(v)
		defer ofile.Close()
		if oerr != nil {
			log.Println("5",oerr)
		}
		data := make([]byte, 10)
		for {
			n, nerr := ofile.Read(data)
			if nerr != nil{
				if nerr == io.EOF{
					break
				}
				log.Println("6",nerr)
			}
			for _,nv := range data[:n]{
				if string(nv) == "\n" {
					num++
				}
			}
		}

	}
	fmt.Println("文件行数: ", num+len(files))
}