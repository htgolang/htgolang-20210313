package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"promagentd/client"
	"promagentd/configs"
	"promagentd/utils"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func Config(cli *client.Client, conf *configs.AgentConfig) error {
	content, err := cli.Config(map[string]interface{}{
		"uuid":    conf.Client.UUID,
		"version": conf.Status.Version,
	})
	if err == nil {
		// 获取到配置
		var response struct {
			Ok     bool
			Reason string
			Result struct {
				Updated bool   `json:"updated"`
				Version int64  `json:"version"`
				Config  string `json:"config"`
			}
		}

		if err := json.Unmarshal([]byte(content), &response); err != nil {
			logrus.WithFields(logrus.Fields{
				"content": content,
				"error":   err.Error(),
			}).Error("error on json unmarshal config text")
		} else if !response.Ok {
			logrus.WithFields(logrus.Fields{
				"error": response.Reason,
			}).Error("error on get config")
		} else if !response.Result.Updated {
			logrus.WithFields(logrus.Fields{
				"version": response.Result.Version,
			}).Info("config not changed")
		} else {
			// 更新
			if err := handle(conf, response.Result.Config); err == nil {
				conf.Status.Version = response.Result.Version
				logrus.WithFields(logrus.Fields{
					"content": response.Result.Config,
					"version": response.Result.Version,
				}).Info("更新配置")
			} else {
				logrus.WithFields(logrus.Fields{
					"content": response.Result.Config,
					"version": response.Result.Version,
					"error":   err,
				}).Error("更新配置")
			}
		}
	}
	return nil
}

type job struct {
	JobName   string `yaml:"job_name"`
	BasicAuth struct {
		Username string `yaml:"username,omitempty"`
		Password string `yaml:"password,omitempty"`
	} `yaml:"basic_auth,omitempty"`
	StaticConfigs []struct {
		Targets []string
	} `yaml:"static_configs"`
}

type prometheusConfig struct {
	Global        interface{} `yaml:"global"`
	Alerting      interface{} `yaml:"alerting"`
	RuleFiles     interface{} `yaml:"rule_files"`
	ScrapeConfigs []job       `yaml:"scrape_configs"`
}

// 处理配置
func handle(conf *configs.AgentConfig, scrapeConfigs string) error {
	fmt.Println(conf.Prometheus)
	// 读取模板配置
	tplContent := utils.ReadFile(conf.Prometheus.TplFile)
	// 生成配置
	var prometheusConfig prometheusConfig
	if err := yaml.Unmarshal([]byte(tplContent), &prometheusConfig); err != nil {
		return err
	}

	var jobs []job
	if err := yaml.Unmarshal([]byte(scrapeConfigs), &jobs); err != nil {
		return err
	}

	prometheusConfig.ScrapeConfigs = append(prometheusConfig.ScrapeConfigs, jobs...)

	config, err := yaml.Marshal(prometheusConfig)
	if err != nil {
		return err
	}

	bakFile := conf.Prometheus.ConfigFile + ".bak"
	utils.WirteFile(bakFile, string(config))

	// 检查配置
	cmd := fmt.Sprintf(`"%s" check config "%s"`, conf.Prometheus.PromTool, bakFile)
	if err := utils.Exec(cmd); err != nil {
		return err
	}

	// 应用配置
	if err := os.Rename(bakFile, conf.Prometheus.ConfigFile); err != nil {
		return err
	}

	// 通知prometheus reload
	if err := utils.Exec(conf.Prometheus.ReloadCmd); err != nil {
		return err
	}

	return nil
}
