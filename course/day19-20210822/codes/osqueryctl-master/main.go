package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/kolide/osquery-go"
	"github.com/kolide/osquery-go/plugin/config"
	"github.com/kolide/osquery-go/plugin/distributed"
	"github.com/kolide/osquery-go/plugin/logger"
	"github.com/kolide/osquery-go/plugin/table"
)

func main() {
	socketPath := flag.String("socket", "", "path to osqueryd extensions socket")
	flag.Int("timeout", 0, "")
	flag.Int("interval", 0, "")
	flag.Parse()

	server, err := osquery.NewExtensionManagerServer("manager", *socketPath)
	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	server.RegisterPlugin(distributed.NewPlugin("distributed", getQueries, writeResults))
	server.RegisterPlugin(logger.NewPlugin("logger", info))
	server.RegisterPlugin(table.NewPlugin("ktable", columns(), body))
	server.RegisterPlugin(config.NewPlugin("config", configs))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func getQueries(ctx context.Context) (*distributed.GetQueriesResult, error) {
	return &distributed.GetQueriesResult{Queries: map[string]string{"time": "select * from time"}}, nil
}

func writeResults(ctx context.Context, results []distributed.Result) error {
	fmt.Println(results)
	return nil
}

func info(ctx context.Context, typ logger.LogType, logText string) error {
	log.Printf("%s: %s\n", typ, logText)
	return nil
}

func columns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.BigIntColumn("id"),
		table.TextColumn("name"),
		table.IntegerColumn("attr"),
		table.DoubleColumn("size"),
	}
}

func body(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	return []map[string]string{
		{
			"id":   "1",
			"name": "kk1.txt",
			"attr": "6",
			"size": "1024",
		}, {
			"id":   "2",
			"name": "kk2.txt",
			"attr": "6",
			"size": "2048",
		},
	}, nil
}

func configs(ctx context.Context) (map[string]string, error) {
	return map[string]string{
		"config": `
{
	"schedule": {
		"users": {
			"query": "SELECT * FROM users;",
			"interval": 60
		}
	},
	"decorators": {
		"load": [
			"SELECT instance_id as id FROM osquery_info;",
			"SELECT uuid AS uuid FROM system_info;"
		],
		"always": [
			"SELECT user AS username FROM logged_in_users WHERE user <> '' ORDER BY time LIMIT 1;"
		]
	},
	"file_paths": {
		"homes": [
			"/root/.ssh/%%",
			"/home/%/.ssh/%%"
		],
		"etc": [
			"/etc/%%"
		],
		"tmp": [
			"/tmp/%%"
		]
	}
}
`,
	}, nil
}
