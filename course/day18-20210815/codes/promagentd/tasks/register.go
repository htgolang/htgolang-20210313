package tasks

import (
	"promagentd/client"
	"promagentd/configs"
	"time"
)

func Regsiter(cli *client.Client, conf *configs.AgentConfig) error {
	cli.Register(map[string]interface{}{
		"uuid":     conf.Client.UUID,
		"addr":     conf.Client.Addr,
		"hostname": conf.Client.Hostname,
		"now":      time.Now().Unix(),
	})

	return nil
}
