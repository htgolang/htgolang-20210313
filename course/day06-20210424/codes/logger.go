package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("web.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Print("第一条日志")
}
