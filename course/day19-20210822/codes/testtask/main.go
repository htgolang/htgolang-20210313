package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/distributed"
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

	mgr, err := osquery.NewExtensionManagerServer("testtask", socket, intervalOption, timeoutOption)
	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.Create("/tmp/task.result.log")
	defer file.Close()

	log.SetOutput(file)

	mgr.RegisterPlugin(distributed.NewPlugin("testtask", func(ctx context.Context) (*distributed.GetQueriesResult, error) {
		log.Printf("query")
		result := new(distributed.GetQueriesResult)
		result.Queries = map[string]string{
			"user_check": "select * from users;",
		}

		return result, nil
	}, func(ctx context.Context, results []distributed.Result) error {
		log.Printf("task: %#v", results)
		return nil
	}))
	if err := mgr.Run(); err != nil {
		log.Fatal(err)
	}
}
