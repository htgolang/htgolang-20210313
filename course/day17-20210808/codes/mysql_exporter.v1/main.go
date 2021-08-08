package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/go-sql-driver/mysql"

	"mysql_exporter/collectors"
	"mysql_exporter/configs"
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

func initMetrics(db *sql.DB) {
	prometheus.MustRegister(collectors.NewHealthyCollector(db))
	prometheus.MustRegister(collectors.NewSlowQueriesCollector(db))
	prometheus.MustRegister(collectors.NewQpsCollector(db))
	prometheus.MustRegister(collectors.NewCommandCollector(db))
	prometheus.MustRegister(collectors.NewConnectionCollector(db))
	prometheus.MustRegister(collectors.NewTrafficCollector(db))
}

func main() {
	// 配置
	// mysql 配置
	// web basic 认证用户名密码
	// 日志

	gconf := configs.ExporterConfig{
		MySQL: configs.MySQLConfig{
			"10.0.0.2",
			3306,
			"golang",
			"golang@2021",
			"cmdb",
		},
		Web: configs.WebConfig{
			":11111",
			struct {
				User     string
				Password string
			}{"", ""},
		},
	}

	// 初始化
	// 初始化日志
	// 初始化数据库
	db, err := initDb(gconf.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 初始化指标
	initMetrics(db)

	http.Handle("/metrics/", promhttp.Handler())

	log.Fatal(http.ListenAndServe(gconf.Web.Addr, nil))
}
