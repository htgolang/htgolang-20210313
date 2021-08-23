package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
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

	mgr, err := osquery.NewExtensionManagerServer("testtable", socket, intervalOption, timeoutOption)
	if err != nil {
		log.Fatal(err)
	}

	mgr.RegisterPlugin(table.NewPlugin("tmp_files", columns(), datas))
	if err := mgr.Run(); err != nil {
		log.Fatal(err)
	}
}

func columns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("file"),
		table.IntegerColumn("size"),
	}
}

func datas(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	infos, err := ioutil.ReadDir("/tmp")
	if err != nil {
		return nil, err
	}
	rt := make([]map[string]string, 0, len(infos))
	for _, info := range infos {
		rt = append(rt, map[string]string{
			"file": info.Name(),
			"size": strconv.FormatInt(info.Size(), 10),
		})
	}
	return rt, nil
}
