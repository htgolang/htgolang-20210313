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

func (s *prometheusAgentService) Query(q string) []*models.PrometheusAgent {
	var agents []*models.PrometheusAgent
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.PrometheusAgent))
	cond := orm.NewCondition()
	cond = cond.And("DeletedAt__isnull", true)
	if q != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("Hostname__icontains", q)
		cond = cond.AndCond(queryCond)
	}
	queryset.SetCond(cond).All(&agents)
	return agents
}

func (s *prometheusAgentService) GetById(id int64) *models.PrometheusAgent {
	ormer := orm.NewOrm()
	agent := &models.PrometheusAgent{Id: id}
	if err := ormer.Read(agent); err == nil {
		return agent
	}
	// 是否已经删除 DeletedAt
	return nil
}

func (s *prometheusAgentService) Modify(agent *models.PrometheusAgent) {
	ormer := orm.NewOrm()
	agent.ConfigVersion = time.Now().Unix()
	ormer.Update(agent, "Config", "ConfigVersion")
}

func (s *prometheusAgentService) Delete(id int64) {
	ormer := orm.NewOrm()
	ormer.QueryTable(new(models.PrometheusAgent)).Filter("Id", id).Update(orm.Params{
		"DeletedAt": time.Now(),
	})
}

func (s *prometheusAgentService) GetConfigByUUID(uuid string, version int64) (string, int64, error) {
	ormer := orm.NewOrm()
	agent := &models.PrometheusAgent{UUID: uuid}
	if err := ormer.Read(agent, "UUID"); err != nil {
		return "", 0, err
	}

	// 是否已经删除 DeletedAt
	// version 更新到数据库

	if agent.ConfigVersion <= version {
		return "", agent.ConfigVersion, nil
	}

	// 拉去配置的时间
	return agent.Config, agent.ConfigVersion, nil
}

type prometheusAlertService struct{}

var PrometheusAlertService = new(prometheusAlertService)

func (s *prometheusAlertService) Trigger(alert *models.PrometheusAlert) {
	/*
	   1. firing
	       数据库中添加记录就行了
	       a cpu 高
	       a cpu 高

	       fingerprint 查询 有没有相同的告警 未解决的
	           有 不用存储
	           无 增加
	   2. resolved:
	       查询 fingerprint 查询 有相同的告警 未解决的
	           有 更新状态解决  endsAt
	*/
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.PrometheusAlert))
	queryset = queryset.Filter("Fingerprint", alert.Fingerprint).
		Filter("Status", models.StatusAlertFiring).
		Filter("DeletedAt__isnull", true)

	if alert.Status == models.StatusAlertFiring {
		if !queryset.Exist() {
			ormer.Insert(alert)
		}
	} else {
		queryset.Update(orm.Params{
			"Status": models.StatusAlertResolved,
			"EndsAt": alert.EndsAt,
		})
	}
}

func (s *prometheusAlertService) Count(q string) int64 {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.PrometheusAlert))
	cond := orm.NewCondition()
	cond.And("DeletedAt__isnull", true)
	if q != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("Instance__icontains", q)
		queryCond = queryCond.Or("Name__icontains", q)
		queryCond = queryCond.Or("Name__icontains", q)
		queryCond = queryCond.Or("Description__icontains", q)

		cond = cond.AndCond(queryCond)
	}

	if total, err := queryset.SetCond(cond).Count(); err != nil {
		return 0
	} else {
		return total
	}

}

func (s *prometheusAlertService) Query(q string, offset int, limit int) []models.PrometheusAlert {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.PrometheusAlert))
	cond := orm.NewCondition()
	cond.And("DeletedAt__isnull", true)
	if q != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("Instance__icontains", q)
		queryCond = queryCond.Or("Name__icontains", q)
		queryCond = queryCond.Or("Name__icontains", q)
		queryCond = queryCond.Or("Description__icontains", q)

		cond = cond.AndCond(queryCond)
	}

	var alerts []models.PrometheusAlert
	queryset.SetCond(cond).OrderBy("-CreatedAt").Limit(limit).Offset(offset).All(&alerts)
	return alerts
}
