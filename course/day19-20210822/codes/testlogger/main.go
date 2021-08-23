package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/logger"
)

func main() {
	var (
		socket   string
		timeout  int
		interval int
	)
	flag.StringVar(&socket, "socket", "", "socket path")
	flag.IntVar(&timeout, "timeout", 3, "connect timeout")
	flag.IntVar(&interval, "interval", 3, "ping interval")

	flag.Usage = func() {
		fmt.Println("testtable --socket socket_path [--timeout 3] [--interval 3]")
		flag.PrintDefaults()
	}

	flag.Parse()

	intervalOption := osquery.ServerPingInterval(time.Duration(interval) * time.Second)
	timeoutOption := osquery.ServerTimeout(time.Duration(timeout) * time.Second)

	mgr, err := osquery.NewExtensionManagerServer("testlogger", socket, intervalOption, timeoutOption)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("/tmp/testosquery.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	mgr.RegisterPlugin(logger.NewPlugin("logger", func(ctx context.Context, typ logger.LogType, record string) error {
		log.Printf("%#v %s", typ, record)
		return nil
	}))

	if err := mgr.Run(); err != nil {
		log.Fatal(err)
	}
}
