package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// HUP
	//
	interrupt := make(chan os.Signal)
	reload := make(chan os.Signal)
	kill := make(chan os.Signal)

	signal.Notify(interrupt, syscall.SIGINT)
	signal.Notify(reload, syscall.SIGHUP)
	signal.Notify(kill, syscall.SIGTERM)

	for {
		select {
		case <-interrupt:
			fmt.Println("interrupt")
		case <-reload:
			fmt.Println("reload")
		case <-kill:
			fmt.Println("kill")
		}
	}

}
