package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Args[1])
	err := exec.Command("bash", "-c", os.Args[1]).Run()
	fmt.Println("err:", err)
}
