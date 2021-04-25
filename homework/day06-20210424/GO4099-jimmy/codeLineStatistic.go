// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"path/filepath"
// )

// /*
//  代码行数统计
//     输入: 目录/文件 文件后缀
//     输出: Go程序的代码行数(空行, }, #)

//     思路：递归，统计\n数量，带缓冲区IO

// */

// func lineStatistic(filename string) {
// 	fmt.Println(filename)
// 	lines := 1
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	for {
// 		_, err := reader.ReadString('\n')
// 		if err != nil {
// 			if err != io.EOF {
// 				fmt.Println(err)
// 				return
// 			}
// 			break
// 		}
// 		lines++
// 	}
// 	fmt.Printf("文件:%v 行数:%d\n", filename, lines)
// }

// func dir(dirname string) {
// 	dirEntrys, err := os.ReadDir(dirname)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, dirEntry := range dirEntrys {
// 		if dirEntry.IsDir() {
// 			dir(filepath.Join(dirname, dirEntry.Name()))
// 		} else {
// 			lineStatistic(filepath.Join(dirname, dirEntry.Name()))
// 		}
// 	}
// }

// func main() {
// 	fmt.Print("请输入目录/文件/文件名后缀: ")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	fileinfo, err := os.Stat(scanner.Text())
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			fmt.Println("您输入的文件不存在")
// 			return
// 		}
// 		fmt.Println(err)
// 	} else {
// 		if fileinfo.IsDir() {
// 			dir(fileinfo.Name())
// 		} else {
// 			lineStatistic(fileinfo.Name())
// 		}
// 	}
// }
