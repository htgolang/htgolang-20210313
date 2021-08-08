package configs

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

type LogConfig struct{}

type ExporterConfig struct {
	MySQL     MySQLConfig
	Web       WebConfig
	LogConfig LogConfig
}
