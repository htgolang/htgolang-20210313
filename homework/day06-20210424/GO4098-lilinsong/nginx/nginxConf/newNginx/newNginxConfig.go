package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type nginxConf struct {
	Upstream0  string `ini:"upstream.server.0"`
	Upstream2  string `ini:"upstream.server.1"`
	Listen     string `ini:"listen"`
	ServerName string `ini:"server_name"`
	Root       string `ini:"root"`
	SslCert    string `ini:"ssl_certificate"`
	SslKey     string `ini:"ssl_certificate_key"`
}

func main() {
	cfg := ini.Empty()

	c := nginxConf{
		Upstream0:  "172.17.1.1",
		Upstream2:  "172.17.1.2",
		Listen:     "8888",
		ServerName: "www.test.com",
		Root:       "/var/www/html",
		SslCert:    "/app/cert/test.pem",
		SslKey:     "/app/cert/test.key",
	}

	err := cfg.Section("nginx").ReflectFrom(&c)
	if err != nil {
		fmt.Println(err)
	}

	err = cfg.SaveTo("./newNginxConfig.ini")
	if err != nil {
		fmt.Println(err)
	}
}
