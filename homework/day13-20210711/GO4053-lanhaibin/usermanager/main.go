package main

import (
	_ "usermanager/routers"

	"usermanager/cli"
)

func main() {
	cli.Cmd.Execute()
}
