package tasks

import (
	"promagentd/client"
	"promagentd/configs"
)

type Task func(cli *client.Client, conf *configs.AgentConfig) error
