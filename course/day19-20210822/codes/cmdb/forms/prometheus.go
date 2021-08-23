package forms

import (
	"cmdb/models"
	"cmdb/services"
	"encoding/json"
	"time"

	"github.com/beego/beego/v2/core/validation"
	"gopkg.in/yaml.v2"
)

type AgentdHeartbeatForm struct {
	UUID       string `json:"uuid"`
	ClientTime int64  `json:"now"`
}

type AgentdRegisterForm struct {
	UUID     string `json:"uuid"`
	Addr     string `json:"addr"`
	Hostname string `json:"hostname"`
	Now      int64  `json:"now"`
}

type AgentdConfigForm struct {
	UUID    string `form:"uuid"`
	Version int64  `form:"version"`
}

type ModifyAgentForm struct {
	Id     int64  `form:"id"`
	Config string `form:"config"`
}

func (f *ModifyAgentForm) Valid(v *validation.Validation) {
	// 验证
	if agent := services.PrometheusAgentService.GetById(f.Id); agent == nil {
		v.SetError("default", "操作对象不存在")
		return
	}

	// 检查数据格式
	var jobs []Job
	if err := yaml.UnmarshalStrict([]byte(f.Config), &jobs); err != nil {
		v.SetError("config", err.Error())
		return
	}
	// 格式化config
	if prettyConfig, err := yaml.Marshal(jobs); err == nil {
		f.Config = string(prettyConfig)
	}
	// 检查数据内容
}

func (f *ModifyAgentForm) ToModel() *models.PrometheusAgent {
	return &models.PrometheusAgent{
		Id:     f.Id,
		Config: f.Config,
	}
}

func (f *ModifyAgentForm) FillModel(agent *models.PrometheusAgent) {
	if agent == nil {
		return
	}
	f.Id = agent.Id

	f.Config = agent.Config

	var jobs []Job
	if f.Config == "" {
		jobs = append(jobs, Job{})
	} else if err := yaml.UnmarshalStrict([]byte(f.Config), &jobs); err != nil {
		return
	}

	// 格式化config
	if prettyConfig, err := yaml.Marshal(jobs); err == nil {
		f.Config = string(prettyConfig)
	}

}

type Job struct {
	JobName   string `yaml:"job_name"`
	BasicAuth struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"basic_auth"`
	StaticConfigs []struct {
		Targets []string `yaml:"targets"`
	} `yaml:"static_configs"`
}

type PrometheusAlertForm struct {
	Fingerprint  string
	Status       string
	GeneratorURL string `json:"generatorURL"`

	StartsAt    *time.Time
	EndsAt      *time.Time
	Labels      map[string]string
	Annotations map[string]string
}

/*
 {
	"status": "firing",
	"labels": {
		"alertname": "Node CPU Percent is High",
		"env": "dev",
		"instance": "localhost:9102",
		"severity": "high"
	},
	"annotations": {
		"description": "Node localhost:9102 CPU Percent is High 200",
		"summary": "Node localhost:9102 CPU Percent is High"
	},
	"startsAt": "2021-08-15T09:20:26.851Z",
	"endsAt": "0001-01-01T00:00:00Z",
	"generatorURL": "http://centos:9090/graph?g0.expr=node_cpu_percent+%3E%3D+80\u0026g0.tab=1",
	"fingerprint": "bd97c3c7d278748a"
}
*/

func (f *PrometheusAlertForm) ToModel() *models.PrometheusAlert {
	var alert models.PrometheusAlert
	alert.Fingerprint = f.Fingerprint
	alert.Instance = f.Labels["instance"]
	alert.Name = f.Labels["alertname"]
	alert.Description = f.Annotations["description"]
	alert.Severity = f.Labels["severity"]
	alert.GeneratorURL = f.GeneratorURL

	alert.StartsAt = f.StartsAt
	if f.Status == "firing" {
		alert.Status = models.StatusAlertFiring
	} else {
		alert.EndsAt = f.EndsAt
		alert.Status = models.StatusAlertResolved
	}

	labels, _ := json.Marshal(f.Labels)
	alert.Labels = string(labels)
	annotations, _ := json.Marshal(f.Annotations)
	alert.Annotations = string(annotations)
	return &alert
}

type PrometheusAlertsForm struct {
	Alerts []PrometheusAlertForm `json:"alerts"`
}
