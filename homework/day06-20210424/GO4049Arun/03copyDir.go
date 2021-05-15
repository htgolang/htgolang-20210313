package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//Judging whether the file is existed
func Exist3(filename string) bool {
	absFilename, _ := filepath.Abs(filename)
	_, err := os.Stat(absFilename)
	if err == nil {
		//if !os.IsNotExist(err) {
		return true
	}
	return false
}

//Judging whether the path is directory
func IsDir3(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CopyFile(srcFile, dstFile string, bufferSize int) {
	//(1)打开源文件
	src, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	//(2)创建目的文件
	dst, err := os.Create(dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()
	data := make([]byte, bufferSize)
	for {
		//从源文件读取内容
		n, err := src.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		//往目的文件写入内容
		n, err = dst.Write(data[:n])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CopyDir(srcDir, dstDir string, bufferSize int) {
	absDirname, _ := filepath.Abs(srcDir)
	if IsDir3(absDirname) {
		files, _ := ioutil.ReadDir(absDirname)
		os.Mkdir(dstDir, os.ModePerm)
		for _, file := range files {
			absPath, _ := filepath.Abs(absDirname + string(os.PathSeparator) + file.Name())
			dstPath, _ := filepath.Abs(dstDir + string(os.PathSeparator) + file.Name())
			if IsDir3(absPath) {
				CopyDir(absPath, dstPath, bufferSize)
			} else {
				CopyFile(absPath, dstPath, bufferSize)
			}
		}
	} else {
		CopyFile(srcDir, dstDir, bufferSize)
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: 03copyDir.exe srcDir|srcFile dstDir|dstFile")
	}
	bufferSize := 10240
	srcDir, dstDir := os.Args[1], os.Args[2]
	if Exist3(srcDir) {
		CopyDir(srcDir, dstDir, bufferSize)
	} else {
		fmt.Println("请检查源文件夹是否存在!")
	}
} /*
(1)文件copy
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note>03copyDir.exe 03copyDir.go 04Test.go
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note>03copyDir.exe 02codeCount.exe 02Test.exe
(2)文件夹copy
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note>03copyDir.exe test test1
*/
