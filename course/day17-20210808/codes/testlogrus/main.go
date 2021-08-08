package main

import (
	"github.com/sirupsen/logrus"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	// handler, err := os.Create("logs/run.log")
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	handler := &lumberjack.Logger{
		Filename:   "logs/run.log",
		MaxSize:    1,
		MaxBackups: 7,
		LocalTime:  false,
		Compress:   false,
	}

	defer handler.Close()

	// logrus.SetLevel(logrus.InfoLevel)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetOutput(handler)

	for i := 0; i <= 10000; i++ {
		logrus.WithFields(logrus.Fields{
			"test":  "1",
			"test2": 2,
		}).Error("error")
		logrus.Warning("warning")
		logrus.Info("info")
		logrus.Debug("debug")
	}
}
