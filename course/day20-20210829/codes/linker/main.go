package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"linker/client"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/imroc/req"

	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/config"
	"github.com/osquery/osquery-go/plugin/distributed"
	"github.com/osquery/osquery-go/plugin/logger"
)

type GConfig struct {
	ID       string
	Hostname string
	Addr     string
	Endpoint string

	ConfigVersion string
}

func initConfig(socketPath string) (*GConfig, error) {
	gconf := new(GConfig)
	client, err := osquery.NewClient(socketPath, 3*time.Second)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	query := func(sql string, key string) (string, error) {
		resp, err := client.Query(sql)
		if err != nil {
			return "", err
		}
		if resp.Status.Code != 0 || len(resp.Response) == 0 {
			return "", fmt.Errorf("query osquery error")
		}
		return resp.Response[0][key], nil

	}

	gconf.ID, err = query("select uuid from osquery_info;", "uuid")
	if err != nil {
		return nil, err
	}

	gconf.Hostname, err = query("select computer_name from system_info;", "computer_name")
	if err != nil {
		return nil, err
	}

	gconf.Addr, err = query("select address from interface_addresses where not(address like '%:%') and not (address like '127.%');", "address")
	if err != nil {
		return nil, err
	}

	// 从配置中读取
	gconf.Endpoint = "http://10.0.0.1:10001"
	return gconf, nil
}

func main() {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.Fatal(err)
	}

	handler := &lumberjack.Logger{
		Filename:   filepath.Join(currentDir, "logs", "linker.log"),
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     17,
		Compress:   true,
	}
	defer handler.Close()
	logrus.SetOutput(handler)

	// 命令行参数(osquery传递)
	var (
		socket   string
		timeout  int
		interval int
		verbose  bool
	)

	flag.StringVar(&socket, "socket", "", "socket path")
	flag.IntVar(&timeout, "timeout", 3, "connect timeout")
	flag.IntVar(&interval, "interval", 3, "ping interval")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose informational messages")

	flag.Usage = func() {
		fmt.Println("linker --socket socket_path [--timeout 3] [--interval 3] [--verbose]")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
		req.Debug = true
	}

	// 定义osquery插件参数
	intervalOption := osquery.ServerPingInterval(time.Duration(interval) * time.Second)
	timeoutOption := osquery.ServerTimeout(time.Duration(timeout) * time.Second)

	gconf, err := initConfig(socket)
	if err != nil {
		logrus.Fatal(err)
	}

	cli := client.NewClient(gconf.ID, gconf.Endpoint)

	// 定义插件管理器
	mgr, err := osquery.NewExtensionManagerServer("linker", socket, intervalOption, timeoutOption)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("linker plugin is running: %d", os.Getpid())
	logrus.Infof("linker config: %v", gconf)

	// 注册插件
	// 配置插件

	mgr.RegisterPlugin(config.NewPlugin("config", func(ctx context.Context) (map[string]string, error) {
		return getConf(cli, gconf, ctx)
	}))
	logrus.Info("Register [config] plugin")

	// 日志记录插件
	mgr.RegisterPlugin(logger.NewPlugin("logger", func(ctx context.Context, typ logger.LogType, record string) error {
		return writeLog(cli, ctx, typ, record)
	}))
	logrus.Info("Register [logger] plugin")

	// 任务插件(获取任务/记录任务结果)
	mgr.RegisterPlugin(distributed.NewPlugin("task", func(ctx context.Context) (*distributed.GetQueriesResult, error) {
		return getTasks(cli, ctx)
	}, func(ctx context.Context, results []distributed.Result) error {
		return taskResults(cli, ctx, results)
	}))
	logrus.Info("Register [task] plugin")

	// 心跳
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			heartbeat(cli)
			<-ticker.C
		}
	}()

	// 注册
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for {
			register(cli, gconf)
			<-ticker.C
		}
	}()

	// 运行插件
	if err := mgr.Run(); err != nil {
		logrus.Fatal(err)
	}
}

func heartbeat(cli *client.Client) {
	logrus.Debug("heartbeat")
	cli.Heartbeat(nil)
}

func register(cli *client.Client, gconf *GConfig) {
	logrus.Debug("register")
	cli.Register(map[string]interface{}{
		"uuid":     gconf.ID,
		"hostname": gconf.Hostname,
		"addr":     gconf.Addr,
	})
}

// 获取配置
func getConf(cli *client.Client, gconf *GConfig, ctx context.Context) (map[string]string, error) {
	var rs struct {
		Ok     bool              `json:"ok"`
		Error  string            `json:"error"`
		Result map[string]string `json:"result"`
	}
	if err := cli.Config(gconf.ConfigVersion, &rs); err != nil {
		logrus.Errorf("Error Get Config: %s", err)
		return nil, err
	}
	if rs.Ok {
		logrus.Debugf("Get Config: %v", rs.Result)
		return rs.Result, nil
	}

	return nil, fmt.Errorf(rs.Error)
}

func writeLog(cli *client.Client, ctx context.Context, typ logger.LogType, record string) error {
	logrus.Debugf("Write Log: [%d]%s", typ, record)
	var rs map[string]interface{}
	cli.Log(rs)
	return nil
}

func getTasks(cli *client.Client, ctx context.Context) (*distributed.GetQueriesResult, error) {
	logrus.Debugf("Get Tasks: %s", "tasks")
	result := new(distributed.GetQueriesResult)
	var rs struct {
		Ok     bool              `json:"ok"`
		Error  string            `json:"error"`
		Result map[string]string `json:"result"`
	}
	if err := cli.Tasks(&rs); err != nil {
		logrus.Errorf("Error Get Tasks: %s", err)
		return nil, err
	}
	if rs.Ok {
		logrus.Debugf("Get Tasks: %v", rs.Result)
		result.Queries = rs.Result
		return result, nil
	}

	return nil, fmt.Errorf(rs.Error)
}

type TaskResult struct {
	Job    int64               `json:"job"`
	Status int                 `json:"status"`
	Result []map[string]string `json:"result"`
}

func taskResults(cli *client.Client, ctx context.Context, results []distributed.Result) error {
	logrus.Debugf("task Resuls: %v", results)

	rs := make(map[string]interface{})
	for _, result := range results {
		job, err := strconv.ParseInt(result.QueryName, 10, 64)
		if err != nil {
			logrus.Errorf("error job[%s] is not int64", result.QueryName)
			continue
		}
		rs[result.QueryName] = &TaskResult{
			Job:    job,
			Status: result.Status,
			Result: result.Rows,
		}
	}
	cli.Results(rs)
	return nil
}
