package configs

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Addr  string
	Token string
}

type PrometheusConfig struct {
	ConfigFile string
}

type ClientConfig struct {
	UUID     string
	Addr     string
	Hostname string
}

type AgentConfig struct {
	Server     ServerConfig
	Prometheus PrometheusConfig
	Client     ClientConfig
}

func getUUID() (string, error) {
	path := "./etc/promagentd.uuid"
	if ctx, err := ioutil.ReadFile(path); err == nil {
		if string(ctx) != "" {
			return string(ctx), nil
		}
	}

	// 无UUID
	agentId := strings.ReplaceAll(uuid.New().String(), "-", "")
	ioutil.WriteFile(path, []byte(agentId), os.ModePerm)
	return agentId, nil
}

func getAddr() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if strings.Index(addr.String(), ":") > 0 {
			continue
		}
		if strings.HasPrefix(addr.String(), "127.") {
			continue
		}

		nodes := strings.SplitN(addr.String(), "/", 2)
		if len(nodes) == 2 {
			return nodes[0], nil
		}
	}
	return "", fmt.Errorf("not ipv4 addr")
}

func Parse() (AgentConfig, error) {
	var agentConfig AgentConfig

	parser := viper.New()
	parser.SetConfigType("yaml")
	parser.SetConfigName("promagentd")

	parser.AddConfigPath(".")
	parser.AddConfigPath("./etc/")

	if err := parser.ReadInConfig(); err != nil {
		return agentConfig, err
	}

	if err := parser.Unmarshal(&agentConfig); err != nil {
		return agentConfig, err
	}

	if agentId, err := getUUID(); err == nil {
		agentConfig.Client.UUID = agentId
	} else {
		return agentConfig, err
	}

	if addr, err := getAddr(); err == nil {
		agentConfig.Client.Addr = addr
	} else {
		return agentConfig, err
	}

	if hostname, err := os.Hostname(); err == nil {
		agentConfig.Client.Hostname = hostname
	} else {
		return agentConfig, nil
	}

	return agentConfig, nil
}