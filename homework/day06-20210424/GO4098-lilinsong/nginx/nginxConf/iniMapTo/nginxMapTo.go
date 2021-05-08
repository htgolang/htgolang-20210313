package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
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
	cfg, err := ini.Load("../nginx.ini")
	fmt.Println(err)


	c1 := new(nginxConf)
	cfg.MapTo(c1)

	cfg.Section("nginx1").Key("listen").SetValue("9999")
	cfg.Section("nginx1").Key("upstream.server.0").SetValue("172.17.22.22")

	cfg.WriteTo(os.Stdout)

}
