package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	force     bool
	recursive bool
)

func init() {
	flag.BoolVar(&force, "f", false, "force to cp")
	flag.BoolVar(&recursive, "r", false, "recursive to cp")
	flag.Parse()
}

func copyFile(srcFile, dstFile string) {
	// 转换为绝对路径
	// fmt.Println(srcFile, dstFile)
	srcFile, err := filepath.Abs(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	dstFile, err = filepath.Abs(dstFile)
	if err != nil {
		log.Fatal(err)
	}

	sr := fileExit(srcFile)
	dr := fileExit(dstFile)
	// 如果-r false,src存在为一个文件
	if !recursive && sr == nil && !isDir(srcFile) {
		// 如果dst 为目录,则dst basename取src basenaem
		if dr == nil && isDir(dstFile) {
			dstFile = filepath.Join(dstFile, getFilename(srcFile))
		}
		// 如果dst 存在,确认是否覆盖
		if _, err := os.Stat(dstFile); err == nil {
			if !force {
				fmt.Printf("%s exited overwrite?[y/n]\n", dstFile)
			}
			var ans string
			fmt.Scan(&ans)
			if ans != "y" {
				return
			}
		}
		copyFileAction(srcFile, dstFile)
	} else if recursive && sr == nil && dr == nil && isDir(srcFile) && isDir(dstFile) {
		// 如果-r true,src dst都是目录
		if _, err := os.Stat(filepath.Join(dstFile, getFilename(srcFile))); err != nil {
			fmt.Println("目录不存在,创建目录！")
			if err := os.Mkdir(filepath.Join(dstFile, getFilename(srcFile)), os.ModePerm); err != nil {
				log.Fatal(err)
			}
		} else {
			if !force {
				fmt.Printf("%s exited overwrite?[y/n]\n", filepath.Join(dstFile, getFilename(srcFile)))
			}
			var ans string
			fmt.Scan(&ans)
			if ans != "y" {
				return
			}
		}
		copyPathAction(srcFile, filepath.Join(dstFile, getFilename(srcFile)))
	} else {
		fmt.Println(srcFile, dstFile)
		log.Fatal("Src or Dst error!")
		return
	}
}

func copyPathAction(srcFile, dstFile string) {
	// fmt.Println(srcFile, dstFile, "start")
	if _, err := os.Stat(dstFile); err != nil {
		fmt.Println("目录不存在,创建目录！")
		if err := os.Mkdir(dstFile, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	fslist, err := ioutil.ReadDir(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	for _, fs := range fslist {
		if fs.IsDir() {
			// 递归调用
			// fmt.Println("path call", filepath.Join(srcFile, fs.Name()), filepath.Join(dstFile, fs.Name()), fs.Name())
			copyPathAction(filepath.Join(srcFile, fs.Name()), filepath.Join(dstFile, fs.Name()))
		} else {
			// fmt.Println("file call", filepath.Join(srcFile, fs.Name()), filepath.Join(dstFile, fs.Name()), fs.Name())
			copyFileAction(filepath.Join(srcFile, fs.Name()), filepath.Join(dstFile, fs.Name()))
		}
	}
}

func copyFileAction(srcFile, dstFile string) {
	// fmt.Println(srcFile, dstFile)
	src, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	bufferSize := 10
	data := make([]byte, bufferSize)
	for {
		n, err := src.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		n, err = dst.Write(data[:n])
		if err != nil {
			log.Fatal(err)
		}
	}

}

// fileExit 判断文件或目录是否存在
func fileExit(filename string) error {
	_, err := os.Stat(filename)
	return err
}

// isDir 判断已经存在的目录或文件是否为一个目录
func isDir(pathname string) bool {
	fs, err := os.Stat(pathname)
	if err != nil {
		log.Fatal(err)
	}
	return fs.IsDir()
}

// getFilename 获取文件名带后缀
func getFilename(filename string) string {
	filenameWithSuffix := path.Base(filename)
	return filenameWithSuffix
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if flag.NArg() < 2 {
		flag.Usage()
		log.Fatal("Error: Need src and dst file or dirctory！")
		return
	}

	srcFile, dstFile := flag.Arg(0), flag.Arg(1)
	copyFile(srcFile, dstFile)
}
