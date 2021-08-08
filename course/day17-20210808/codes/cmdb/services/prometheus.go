package services

import (
	"time"

	"github.com/beego/beego/v2/client/orm"

	"cmdb/models"
)

type prometheusAgentService struct{}

var PrometheusAgentService = new(prometheusAgentService)

func (s *prometheusAgentService) Heartbeat(uuid string) error {
	ormer := orm.NewOrm()
	agent := models.PrometheusAgent{
		UUID: uuid,
	}
	now := time.Now()
	if err := ormer.Read(&agent, "UUID"); err == nil {
		agent.HeartbeatTime = &now
		agent.DeletedAt = nil
		ormer.Update(&agent, "HeartbeatTime", "DeletedAt")
	}
	return nil
}

func (s *prometheusAgentService) Register(uuid string, addr string, hostname string) error {
	ormer := orm.NewOrm()
	agent := models.PrometheusAgent{
		UUID: uuid,
	}
	now := time.Now()
	if err := ormer.Read(&agent, "UUID"); err == nil {
		agent.Addr = addr
		agent.Hostname = hostname
		agent.HeartbeatTime = &now
		agent.DeletedAt = nil
		ormer.Update(&agent, "Addr", "Hostname", "HeartbeatTime", "DeletedAt")
	} else {
		agent.Addr = addr
		agent.Hostname = hostname
		agent.HeartbeatTime = &now
		ormer.Insert(&agent)
	}
	return nil
}

func (s *prometheusAgentService) GetConfig(uuid string, version int64) error {
	return nil
}
