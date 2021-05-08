package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func countfile(file4 string){

	file, err := os.Open(file4)
	var flag =0
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n') //注意是字符
		flag++
		if err == io.EOF {

			fmt.Println(flag)

			fmt.Println("文件读完了")
			break
		}

	}

}

func countdir(file3 string){

	dirs,err :=ioutil.ReadDir(file3)
	if err ==nil {
		for _, fi := range dirs {
			if fi.IsDir() {
				countdir(filepath.Join(file3,fi.Name()))

			}else {

				countfile(filepath.Join(file3,fi.Name()))
			}
		}
	}

}

func main(){

	file5 :=flag.String("s","","src file")

	if file,err :=os.Stat(*file5);err !=nil{

		panic(err)
		if os.IsNotExist(err){

			fmt.Println("源文件不存在！！！")
		}else{

			fmt.Println("源文件获取错误！！！")
		}
	}else{
		if file.IsDir(){
			countdir(*file5)
		}else {
			countfile(*file5)
		}
	}

}
