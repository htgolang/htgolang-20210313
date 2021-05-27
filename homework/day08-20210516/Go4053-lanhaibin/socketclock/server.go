package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const MAX_CONNECTION = 2

func main() {
	addr := "0.0.0.0:9000"
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}

	stop := make(chan os.Signal)
	quit := make(chan bool)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)

	connections := make(chan net.Conn, MAX_CONNECTION)
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				select {
				case <-quit:
					return
				default:
					log.Println(err)
				}
			}
			connections <- conn
			wg.Add(1)
			go func() {
				defer wg.Done()
				time.Sleep(3 * time.Second)
				conn := <-connections
				fmt.Fprint(conn, time.Now().Format("2006-01-2 15:04:05"))
				conn.Close()
			}()
		}
	}()
	<-stop
	close(quit)
	wg.Wait()
}
