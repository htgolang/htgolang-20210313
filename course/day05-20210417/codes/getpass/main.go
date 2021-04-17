package main

import (
	"fmt"

	"github.com/howeyc/gopass"
	"github.com/princebot/getpass"
)

func main() {
	data, _ := getpass.Get("prompt to display to user goes here:")
	// if err != nil {

	// }
	// []byte byte => 整数
	fmt.Printf("%v, %q", data, string(data))
	data, _ = gopass.GetPasswd()
	fmt.Printf("%v, %q", data, string(data))
}
