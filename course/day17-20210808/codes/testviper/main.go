package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Auth struct {
	Type     string
	User     string
	Password string
}

type webConfig struct {
	Addr string
	Auth Auth
}

type mysqlConfig struct {
	Addr string `mapstructure:"host"`
	Port int
}

type Config struct {
	Web   webConfig
	MySQL mysqlConfig
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("testviper")

	viper.SetDefault("web.addr", ":1000")
	viper.SetDefault("mysql.host", "127.0.0.1")
	viper.SetDefault("mysql.port", 3306)
	viper.SetDefault("mysql.user", "root")
	viper.SetDefault("mysql.password", "")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./etc/")
	viper.AddConfigPath(".")

	// SetConfigPath

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("warning, config not found, use default")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println(viper.GetString("web.addr"))
	fmt.Println(viper.GetString("web.auth.user"))
	fmt.Println(viper.GetString("web.auth.password"))
	fmt.Println(viper.GetString("mysql.host"))
	fmt.Println(viper.GetString("mysql.port"))

	var config Config

	err := viper.Unmarshal(&config)
	fmt.Println(err)
	fmt.Printf("%#v\n", config)

	config.Web.Addr = "1.1.1.1:999"

	viper.Set("web.addr", "1.1.1.1:9999")

	viper.WriteConfigAs("./config.yaml")

}
