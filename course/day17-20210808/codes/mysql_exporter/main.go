package main

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"mysql_exporter/collectors"
	"mysql_exporter/configs"
	"mysql_exporter/handlers"
)

func initDb(config configs.MySQLConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Db,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func initLogger(config configs.LogConfig) func() {
	logger := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		LocalTime:  config.LocalTime,
		Compress:   config.Compress,
	}
	logrus.SetOutput(logger)
	logrus.SetLevel(logrus.Level(config.Level))
	logrus.SetReportCaller(config.ReportCaller)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	return func() {
		logger.Close()
	}
}

func initMetrics(db *sql.DB) {
	prometheus.MustRegister(collectors.NewHealthyCollector(db))
	prometheus.MustRegister(collectors.NewSlowQueriesCollector(db))
	prometheus.MustRegister(collectors.NewQpsCollector(db))
	prometheus.MustRegister(collectors.NewCommandCollector(db))
	prometheus.MustRegister(collectors.NewConnectionCollector(db))
	prometheus.MustRegister(collectors.NewTrafficCollector(db))
}

func main() {
	gconf, err := configs.Parse("./etc/config.yaml")
	if err != nil {
		logrus.Fatal(err)
	}

	// 初始化
	// 初始化日志
	close := initLogger(gconf.Log)
	defer close()

	// 初始化数据库
	db, err := initDb(gconf.MySQL)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	// 初始化指标
	initMetrics(db)

	http.Handle("/metrics/", handlers.AuthHandler(promhttp.Handler(), gconf.Web.Auth.User, gconf.Web.Auth.Password))

	logrus.Fatal(http.ListenAndServe(gconf.Web.Addr, nil))
}
