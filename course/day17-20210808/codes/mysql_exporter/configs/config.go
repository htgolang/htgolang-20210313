package configs

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

type WebConfig struct {
	Addr string
	Auth struct {
		User     string
		Password string
	}
}

type LogConfig struct {
	Filename     string
	MaxSize      int
	MaxBackups   int
	LocalTime    bool
	Compress     bool
	Level        int
	ReportCaller bool
}

type ExporterConfig struct {
	MySQL MySQLConfig
	Web   WebConfig
	Log   LogConfig
}

func setDefault(parser *viper.Viper) {
	parser.SetDefault("web.addr", "0.0.0.0:9999")
	parser.SetDefault("mysql.host", "127.0.0.1")
	parser.SetDefault("mysql.port", 3306)
	parser.SetDefault("log.filename", "logs/run.log")
	parser.SetDefault("log.max_size", 10)
	parser.SetDefault("log.max_backups", 7)
	parser.SetDefault("log.local_time", true)
	parser.SetDefault("log.level", int(logrus.InfoLevel))

}

func Parse(path string) (ExporterConfig, error) {
	var config ExporterConfig

	parser := viper.New()
	setDefault(parser)

	if _, err := os.Stat(path); err == nil {
		parser.SetConfigFile(path)
		if err := parser.ReadInConfig(); err != nil {
			return config, err
		}
	} else if err != nil && os.IsNotExist(err) {
		logrus.Warn("config file is not found, use default config")
	}

	if err := parser.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
