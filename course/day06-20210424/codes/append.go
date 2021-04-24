package main

import (
	"log"
	"os"
	"time"
)

func main() {
	path := "web.log"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(time.Now().Format("2006-01-02 15:04:05\n"))
}
