package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// ioutil
	// io

	data, _ := ioutil.ReadFile("rdwr.go")
	fmt.Println(string(data))

	files, _ := ioutil.ReadDir(".")
	for _, info := range files {
		fmt.Println(info.Name())
	}

	ioutil.WriteFile("ioutil.txt", []byte("xxxxx"), os.ModePerm)

	//

	path, _ := ioutil.TempDir(".", "tmp_")
	fmt.Println(path)
	file, _ := ioutil.TempFile(".", "tmpf_")
	defer file.Close()
	file.WriteString("xxx")

	f1, _ := os.Create("f1.txt")
	f2, _ := os.Create("f2.txt")

	m := io.MultiWriter(f1, f2) // 输出流(可以写数据的) 给所有的对象中写入数据
	m.Write([]byte("abc"))

	f1, _ = os.Open("std.go")
	f2, _ = os.Open("copyfile.go")
	r := io.MultiReader(f1, f2) // 输入流(可以读取数据的) // 按照顺序读取所有文件中的内容

	io.Copy(os.Stdout, r)
}
