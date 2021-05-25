package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Directory exist check
func DirExist(path string) bool {
	stat, err := os.Lstat(path)
	if stat.IsDir() {
		return false
	}
	return !os.IsNotExist(err)
}

// File exist check
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// Help and get input data
func InputPath() (srcFile, dstFile string, help bool) {
	flag.StringVar(&srcFile, "s", "", "Please enter the source file.")
	flag.StringVar(&dstFile, "d", "", "Plsase enter the target file.")
	flag.BoolVar(&help, "h", false, "Help")

	flag.Usage = func() {
		fmt.Printf("Usage: %v -s file -d /path/file \n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	return
}

// Copy srcfile to dstfile
func copyFile(src, dst string) {
	srcFile, err := os.Open(src)
	defer srcFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	dstFile, err := os.Create(dst)
	defer dstFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(dstFile, srcFile)
}

// main
func main() {
	// Get input data
	src, dst, help := InputPath()
	// RUN
	if help || src == "" || dst == "" {
		// Enter information exception to return help information
		flag.Usage()
		return
	} else if !DirExist(src) || !DirExist(dst) {
		// Check that the srcFile and dstFile entered are directories
		fmt.Printf("srcFile:%v or dstFile:%v entered are directory.\n", src, dst)
		return
	} else if !FileExist(src) {
		// Check in the file exist
		fmt.Printf("srcFile:%v or dstFile:%v file does not exist.\n", src, dst)
		return
	} else if FileExist(src) {
		// Thc file already exist
		var opt string
		fmt.Println("The file already exist,Do you want to overwrite the file？(y/n)：")
		fmt.Scan(&opt)
		switch strings.ToLower(opt) {
		case "y", "yes":
			copyFile(src, dst)
		case "n", "no":
			fmt.Println("Bay~")
			return
		default:
			fmt.Println("Command error, exit。Bay~")
			return
		}
	} else {
		// Copy srcFile to dstFile
		copyFile(src, dst)
		return
	}
}
