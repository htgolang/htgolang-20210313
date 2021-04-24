package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.WriteString("a")
	builder.WriteString("1")
	builder.WriteString("a")
	builder.WriteString("2")
	builder.WriteString("a")

	fmt.Println(builder.String())
	builder.Reset()
	fmt.Println(builder.String())

}
