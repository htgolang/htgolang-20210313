package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	directory string
	filename  string
	suffix    string
	h         bool
	version   bool
)

func usage() {
	fmt.Fprintf(os.Stdout, `codeCount version: codeCount/1.0.0
Usage: go build 02codeCount.go && 02codeCount.exe [-hv] [-d directory] [-f filename] [-s suffix]

Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&h, "h/-help", false, "show help")
	flag.BoolVar(&version, "v", false, "show version and exit")
	flag.StringVar(&suffix, "s", "*.go", "set the suffix of file")
	flag.StringVar(&directory, "d", "example", "set the directory name of file")
	flag.StringVar(&filename, "f", "example.go", "set filename")
	flag.Usage = usage //改变默认的Usage
}

//Judging whether the file is existed
func Exist(filename string) bool {
	absFilename, _ := filepath.Abs(filename)
	_, err := os.Stat(absFilename)
	if err == nil {
		//if !os.IsNotExist(err) {
		return true
	}
	return false
}

//Judging whether the path is directory
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

//Judging whether the path is file
func IsFile(path string) bool {
	return !IsDir(path)
}

func FileCount(filename string) {
	if !Exist(filename) {
		fmt.Println("Please check whether the path of file is existed!")
	}
	var blankLine, leftBrace, rightBrace, poundNum, goComment int
	content, err := ioutil.ReadFile(filename)
	if err == nil {
		for i := 0; i < len(content); i++ {
			if 10 == content[i] {
				blankLine++
			} else if 123 == content[i] {
				leftBrace++
			} else if 125 == content[i] {
				rightBrace++
			} else if 35 == content[i] {
				poundNum++
			} else if 47 == content[i] {
				goComment++
			}
		}
	}
	fmt.Printf("%s总行数:%d,左大括号:%d,右大括号:%d,#注释:%d,Golang注释:%d行\n", filename, blankLine+1, leftBrace,
		rightBrace, poundNum, goComment/2)
}

func dirCount(dirName string) {
	absDirname, _ := filepath.Abs(dirName)
	if IsDir(absDirname) {
		files, _ := ioutil.ReadDir(absDirname)
		for _, file := range files {
			absPath, _ := filepath.Abs(absDirname + string(os.PathSeparator) + file.Name())
			if IsDir(absPath) {
				dirCount(absPath)
			} else {
				FileCount(absPath)
			}
		}
	} else {
		FileCount(absDirname)
	}
}

func suffixCount(suffix string) {
	files, err := filepath.Glob(suffix)
	if err == nil {
		for _, file := range files {
			absDirname, _ := filepath.Abs(file)
			FileCount(absDirname)
		}
	}
}
func main() {
	if len(os.Args) < 2 {
		flag.Usage()
	}
	flag.Parse()
	if version {
		fmt.Println("codeCount/1.0.0")
	}
	if len(filename) > 0 && "example.go" != filename {
		if Exist(filename) {
			FileCount(filename)
		} else {
			fmt.Println("路径不存在或错误,请检查!")
		}
	}
	if len(directory) > 0 && "example" != directory {
		if Exist(directory) {
			dirCount(directory)
		} else {
			fmt.Println("路径不存在或错误,请检查!")
		}
	}
	if len(suffix) > 0 {
		suffixCount(suffix)
	}
}

/*

E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note>02codeCount.exe -h
codeCount version: codeCount/1.0.0
Usage: go build 02codeCount.go && 02codeCount.exe [-hv] [-d directory] [-f filename] [-s suffix]

Options:
  -d string
        set the directory name of file (default "test")
  -f string
        set filename (default "test1Going.go")
  -h/-help
        show help
  -s string
        set the suffix of file (default "*.go")
  -v    show version and exit

(1)文件测试
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note>02codeCount.exe -f test/test.go
test/test.go总行数:18,左大括号:1,右大括号:1,#注释:0,Golang注释:2行

(2)目录测试
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note>02codeCount.exe -d test
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note\test\te.c总行数:6,左大括号:0,右大括号:0,#注释:0,Golang注释:0行
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note\test\test.go总行数:18,左大括号:1,右大括号:1,#注释:0,Golang注释:2行
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note\test\test1\test1Children.c总行数:6,左大括号:0,右大括号:0,#注释:2,Golang注释:0行
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note\test\test1\test1Going.go总行数:8,左大括号:1,右大括号:1,#注释:0,Golang注释:0行

(3)文件后缀测试
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note>02codeCount.exe -s test/*.c
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note\test\te.c总行数:6,左大括号:0,右大括号:0,#注释:0,Golang注释:0行

E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note>02codeCount.exe -s test/test1/*.c
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note\test\test1\test1Children.c总行数:6,左大括号:0,右大括号:0,#注释:2,Golang注释:0行

E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note>02codeCount.exe -s *.go
E:\002golang\01NoteAll\20210424day06\02practice\Teacher\note\02codeCount.go总行数:181,左大括号:36,右大括号:35,#注释:9,Golang注释:12行

*/
