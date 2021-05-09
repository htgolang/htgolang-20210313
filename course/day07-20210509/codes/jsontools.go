package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	txt := `[1, 2, 3`
	ok := json.Valid([]byte(txt))
	fmt.Println(ok)
}
