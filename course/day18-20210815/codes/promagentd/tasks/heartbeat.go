package tasks

import (
	"promagentd/client"
	"promagentd/configs"
	"time"
)

func Heartbeat(cli *client.Client, conf *configs.AgentConfig) error {
	// 心跳
	cli.Heartbeat(map[string]interface{}{
		"uuid": conf.Client.UUID,
		"now":  time.Now().Unix(),
	})

	return nil
}
