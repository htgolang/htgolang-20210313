package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var dirSlice []string
var fileSlice []string

//List the dirs and files
func listDirs(srcdir string) {
	files, _ := ioutil.ReadDir(srcdir)
	for _, file := range files {
		//Define the new dir
		filename := srcdir + "/" + file.Name()
		if file.IsDir() {
			listDirs(filename)
			//Append the dirs to the slice
			dirSlice = append(dirSlice, filename)
		} else {
			// Append the files to the slice
			fileSlice = append(fileSlice, filename)
		}
	}
}

func copyContent(dstdir string) {
	//Create the dirs
	for _, dir := range dirSlice {
		//Define the new dirs
		newdir := dstdir + "/" + dir
		os.MkdirAll(newdir, os.ModePerm)
	}

	//Copy the files under the dir
	for _, file := range fileSlice {
		//Open the src files
		src, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer src.Close()

		//Open the dst files
		//Define the new files
		newfile := dstdir + "/" + file
		dst, err := os.Create(newfile)
		if err != nil {
			log.Fatal(err)
		}
		defer dst.Close()

		//Read the data from the src file
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
			//Write the data to the dst file
			n, err = dst.Write(data[:n])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

//Check the dir exist or not
func dirExistNot(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func main() {
	//Define the source directory
	srcdir := "./tmp_532133273"
	if !dirExistNot(srcdir) {
		fmt.Println("==> The source dirname is not exist.")
		return
	}
	listDirs(srcdir)
	//Define the target directory.
	dstdir := "copy-test-v1"
	copyContent(dstdir)
}
