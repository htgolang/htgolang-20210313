package main

import (
	"copydir/copy"
	"log"
)

func main() {
	err := copy.CopyDir("/tmp/foo", "/tmp/bar")
	if err != nil {
		log.Fatal(err)
	}
}
