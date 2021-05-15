package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)
/*
有一个最后输出几行顺序的问题,等看完链表后再来看这个问题,暂时先放这里!!!
 */
func backwardN(filename, lines string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	tailN, _ := strconv.Atoi(lines)
	linesSlice := make([]string, tailN)
	scanner := bufio.NewScanner(file)
	lineCount := 0
	//dividing two cases: headP < tailP and headP > tailP
	headP :=0
	tailP :=0
	//pValidHead :=0
	//pValidTail :=0
	//pTail = pHead + tailN
	for scanner.Scan() {
		for i := 0; i < len(linesSlice); i++ {
			if i == lineCount%tailN {
				//if lineCount>pTail{
				//	headP = headP+1
				//	tailP = tailP+1
				//
				//	tailP = tailP%tailN
				//	headP = headP%tailN
				//}
				tailP = tailP+i
				if tailP +1 > tailN{
					tailP = headP
					headP = headP+i
				}
				linesSlice[i] = scanner.Text()
			}
		}
		lineCount = lineCount + 1
	}

	if headP < tailP{
		for i := 0; i < len(linesSlice); i++ {
			fmt.Println(linesSlice[i])
		}
	}else {
		for i := 0; i < tailP; i++ {
			fmt.Println(linesSlice[i])
		}
		for i := tailP; i < len(linesSlice); i++ {
			fmt.Println(linesSlice[i])
		}
	}
}

func watchFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(0, 2)
	reader := bufio.NewReaderSize(file, 20)
	line := make([]byte, 0, 4096)
	for {
		data, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Print(err)
				break
			}
		} else if isPrefix == true {
			line = append(line, data...)
		} else {

			if len(line) > 0 {

				fmt.Println(string(append(line, data...)))
				line = make([]byte, 0, 4096)
			} else {
				fmt.Println(string(data))
			}
		}
	}
}
func main() {
	if len(os.Args) != 4 {
		log.Fatal("usage: 04tailn.exe -n lines filename")
	}
	lines, filename := os.Args[2], os.Args[3]
	backwardN(filename, lines)
	watchFile(filename)
}
