package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"promagentd/client"
	"promagentd/configs"

	"github.com/sirupsen/logrus"

	"github.com/imroc/req"
)

func main() {
	req.Debug = true
	conf, err := configs.Parse()
	if err != nil {
		logrus.Fatal(err)
	}

	cli := client.NewClient(conf)

	// 心跳 10s
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			// 心跳
			cli.Heartbeat(map[string]interface{}{
				"uuid": conf.Client.UUID,
				"now":  time.Now().Unix(),
			})
			<-ticker.C
		}
	}()

	// 注册
	go func() {
		ticker := time.NewTicker(15 * time.Minute)
		for {
			cli.Register(map[string]interface{}{
				"uuid":     conf.Client.UUID,
				"addr":     conf.Client.Addr,
				"hostname": conf.Client.Hostname,
				"now":      time.Now().Unix(),
			})
			<-ticker.C
		}
	}()

	// 获取
	go func() {
		version := ""
		ticker := time.NewTicker(10 * time.Second)
		for {
			cli.Config(map[string]interface{}{
				"uuid":    conf.Client.UUID,
				"version": version,
			})
			<-ticker.C
		}
	}()

	// 等待
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
}
