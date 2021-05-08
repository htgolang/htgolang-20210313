package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func copyfile(src1,dst1 string){

	content, err := ioutil.ReadFile(src1)
	if err != nil {
		fmt.Println("read file failed, err:", err)
	}
	conts :=string(content)

	file, err := os.OpenFile(dst1, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	//写入字节切片数据
	file.WriteString(conts) //直接写入字符串数据

}

func copydir(src2,dst2 string){

	_, err := os.Stat(dst2)
	if err != nil {
		if os.IsNotExist(err){
			os.Mkdir(dst2,os.ModePerm)
		}
	}
	dirs,err :=ioutil.ReadDir(src2)
	if err ==nil {
		for _, fi := range dirs {
			if fi.IsDir() {
				copydir(filepath.Join(src2,fi.Name()), filepath.Join(dst2,fi.Name()))
				fmt.Println(fi.Name())
			}else {

				copyfile(filepath.Join(src2,fi.Name()), filepath.Join(dst2,fi.Name()))
			}
		}
	}

}

func main(){

	src :=flag.String("s","","src file")
	dst :=flag.String("d","","dst file")
	help :=flag.Bool("h",false,"help")

	flag.Usage=func(){
		fmt.Println(`
Usage: copyfile -s srcfile -d dstfile
Option:
       `)
		flag.PrintDefaults()
	}
	flag.Parse()
	if *help || *src == "" || *dst ==""{
		flag.Usage()
	}else{
		//判断src是否存在，不存在退出，存在如果是文件夹copydir，如果是文件copyfile
		if file,err :=os.Stat(*src);err !=nil{
			if os.IsNotExist(err){
				fmt.Println("源文件不存在！！！")
			}else{
				fmt.Println("源文件获取错误！！！")
			}
		}else{
			if file.IsDir(){
				copydir(*src,*dst)
			}else {
				copyfile(*src,*dst)
			}
		}
	}
}
