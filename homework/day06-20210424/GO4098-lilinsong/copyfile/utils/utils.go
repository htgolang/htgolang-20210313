package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CopyFile struct {
}

func (c *CopyFile) Start(src, dest string) {
	if fileInfo, err := os.Stat(src); os.IsNotExist(err) {
		log.Fatalln(err)
	} else {
		if fileInfo.IsDir() {
			c.srcIsDir(src, dest)
		} else {
			c.srcIsFile(src, dest)
		}
	}
}

func (c CopyFile) srcIsDir(src, dest string) {
	if c.isExist(dest) {
		os.Mkdir(dest, 0755)
		c.copyDir(src, dest)
	}

	file, err := os.Stat(dest)
	if err != nil {
		log.Fatalln(err)
	}

	if !file.IsDir() {
		fmt.Println("src is dir, dest is file")
		return
	}
	c.copyDir(src, dest)
}

func (c CopyFile) isExist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return true
	}
	return false
}

func (c CopyFile) srcIsFile(src, dest string) {
	file, err := os.Open(dest)
	if err != nil {
		if os.IsNotExist(err) {
			c.copyfile(src, dest)
			return
		}
		fmt.Println("src is file", err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("sourceIsFile", err)
		return
	}
	if fileinfo.IsDir() {
		//dest是目录
		if !c.overWrite(src, dest) { //dest目录下有相同名字的文件
			c.copyfile(src, filepath.Join(dest, filepath.Base(src)))
		}
	} else { //文件存在，有相同的名字  dest下有相同名字的文件
		if !c.overWrite(src, dest) {
			c.copyfile(src, dest)
		}

	}
}


func (c *CopyFile) copyDir(src, dest string) {

	fileInfo, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range fileInfo {
		if file.IsDir() {
			if err := os.Mkdir(filepath.Join(dest, file.Name()), 0755); err != nil {
				fmt.Printf("该目录 %s/%s 以存在是否继续? ", dest, file.Name())
				prompt := c.input("输入任意键继续，输入n退出: ")
				if strings.ToLower(prompt) == "n" {
					return
				}
			}

			c.copyDir(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))

		} else {
			if !c.overWrite(filepath.Join(src, file.Name()), dest) {
				c.copyfile(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
			}

		}
	}
}

func (c *CopyFile) input(prompt string) string {
	fmt.Printf(prompt)
	read := bufio.NewReader(os.Stdin)
	if txt, err := read.ReadString('\n'); err != nil {
		fmt.Println(err)
		return ""
	} else {
		return strings.TrimSpace(txt)
	}
}

func (c *CopyFile) overWrite(src, dest string) bool {
	return c.overWriteBase(src, dest, src)
}

func (c *CopyFile) overWriteBase(src, dest, cmd string) bool {
	var overwritePrompt string
	var readDirDest string
	info, _ := os.Stat(dest)
	if info.IsDir() {
		readDirDest = dest
		//dest是目录
		overwritePrompt = filepath.Join(dest, filepath.Base(src))
	} else {
		//路径存在是文件
		readDirDest = filepath.Dir(dest)
		overwritePrompt = dest
	}

	fileInfo, _ := ioutil.ReadDir(readDirDest)
	for _, file := range fileInfo {
		if filepath.Base(cmd) == file.Name() {
			fmt.Printf("cpfile overwrite %s?", overwritePrompt)
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				if strings.ToLower(strings.TrimSpace(scanner.Text())) == "y" {
					var newdest string
					if info.IsDir() {
						newdest = filepath.Join(dest, filepath.Base(src))
					} else {
						newdest = dest
					}
					c.copyfile(src, newdest)
					return true
				} else {
					return true
				}
			}
		}
	}
	return false
}

func (c CopyFile) copyfile(src, dest string) {
	sfile, _ := os.Open(src)
	dfile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sfile.Close()
	defer dfile.Close()

	bytes := make([]byte, 1024*1024*50)
	_, cberr := io.CopyBuffer(dfile, sfile, bytes)
	if cberr != nil {
		fmt.Println(err)
		return
	}
}
