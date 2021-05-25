// 3. copyfile => copydir
//     思路：递归

// 当前开发在windows下，如在linux下运行需要改变目录结构方式

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func copyDir(src, dest string) (bool, error){
	//列出当前目录下的所有
	readinfo,err := ioutil.ReadDir(src)
	if err != nil{
		return false,err
	} 
	//遍历循环主体
	for _,v := range readinfo{
		// 判断当前源是否为目录；是进入递归遍历，并创建这层目录
		if v.IsDir(){
			_, verr := copyDir(src+"\\"+v.Name(), dest)
			if verr != nil{
				return false, verr
			}
			mverr := os.MkdirAll(dest+"\\"+src, os.ModePerm)
			if mverr != nil{
				return false,mverr
			}
		} else { // 为文件时，创建父目录和子文件
			serr := os.MkdirAll(dest+"\\"+src, os.ModePerm)
			if serr != nil{
				return false,serr
			}
			file, ferr := os.Create(dest+"\\"+ src+"\\"+v.Name())
			if ferr != nil{
				return false,ferr
			}
			data := make([]byte, 10)
			vread, verr := os.Open(src+"\\"+v.Name())
			defer vread.Close()
			if verr != nil{
				return false,verr
			}
			for {
				n, nerr := vread.Read(data)
				if nerr != nil{
					if nerr == io.EOF{
						break
					}
					return false,verr
				}
				file.Write(data[:n])
			}
		}
	}
	return true,nil
}

func main(){

	//判断源和目标是否有效
	if len(os.Args) != 3{
		log.Fatalln("usage: copydir.go src dest")
	}
	_,oerr := os.Stat(os.Args[1])
	if os.IsNotExist(oerr){
		log.Fatalln(oerr)
	}
	_,terr := os.Stat(os.Args[2])
	if os.IsNotExist(terr){
		log.Fatalln(terr)
	}

	b, err := copyDir(os.Args[1], os.Args[2])
	if b{
		fmt.Println("success")
	} else{
		fmt.Println(err)
	}
}