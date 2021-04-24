package ioutlis

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(prompt string) string {
	fmt.Printf(prompt)
	read := bufio.NewReader(os.Stdin)
	if txt, err := read.ReadString('\n'); err != nil {
		fmt.Println(err)
		return ""
	} else {
		return strings.TrimSpace(txt)
	}
}
