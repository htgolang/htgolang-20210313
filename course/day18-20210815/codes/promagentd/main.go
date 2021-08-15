package main

import (
	"time"

	"promagentd/client"
	"promagentd/configs"
	"promagentd/runner"
	"promagentd/tasks"

	"github.com/sirupsen/logrus"

	"github.com/imroc/req"
)

func main() {
	req.Debug = false
	logrus.SetLevel(logrus.InfoLevel)
	conf, err := configs.Parse()
	if err != nil {
		logrus.Fatal(err)
	}

	cli := client.NewClient(conf)
	runnerCli := runner.New()

	createRunner := func(interval time.Duration, task tasks.Task) runner.Callback {
		return func() {
			ticker := time.NewTicker(interval)
			for {
				if err := task(cli, conf); err != nil {
					break
				}
				<-ticker.C
			}
		}
	}

	// 心跳 10s
	runnerCli.Run(createRunner(10*time.Second, tasks.Heartbeat))
	// 注册
	runnerCli.Run(createRunner(15*time.Minute, tasks.Regsiter))
	// 获取
	runnerCli.Run(createRunner(10*time.Second, tasks.Config))

	runnerCli.Wait()

}
