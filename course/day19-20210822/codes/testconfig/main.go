package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/config"
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

	mgr, err := osquery.NewExtensionManagerServer("testconfig", socket, intervalOption, timeoutOption)
	if err != nil {
		log.Fatal(err)
	}

	mgr.RegisterPlugin(config.NewPlugin("config", getConf))
	if err := mgr.Run(); err != nil {
		log.Fatal(err)
	}
}

func getConf(ctx context.Context) (map[string]string, error) {
	return map[string]string{
		"config": `
		{
			"schedule": {
				"check_etc_host": {
					"query": "select * from etc_hosts;",
					"interval" : 10
				}
			}
		}
		`,
	}, nil
}
