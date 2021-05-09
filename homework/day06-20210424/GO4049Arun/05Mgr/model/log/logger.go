package log

import (
	"bufio"
	"log"
	"mgr/conf"
	"os"
)

func LogRecord(content string)  {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	//fmt.Println(conf.Info.LogFilepath)
	file,_ := os.OpenFile(conf.Info.LogFilepath,os.O_APPEND|os.O_CREATE,os.ModePerm)
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	log.SetOutput(writer)
	log.Println(content)
}