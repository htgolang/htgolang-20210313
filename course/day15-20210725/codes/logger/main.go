package main

import (
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	// logs.SetLogger(logs.AdapterConsole, `{"level":5}`)
	// logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, `
	{
		"filename" : "web.log",
		"level" : 6,
		"maxlines" : 10000
	}
	`)

	for i := 0; i < 10000; i++ {
		logs.Error("error")
		logs.Warn("warn")
		logs.Info("info")
		logs.Debug("debug")
	}
}
