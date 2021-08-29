package services

import (
	"cmdb/models"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

const (
	JobStatusNew     = 0
	JobStatusRunning = 1
	JobStatusSuccess = 2
	JobStatusFailure = 3
)

var types = []*models.OsqueryTaskType{
	{
		Key: "users", Name: "获取用户信息", SQL: "select * from users;",
		Columns: [][2]string{
			{"username", "用户名"},
			{"description", "描述"},
			{"directory", "主目录"},
			{"shell", "shell"},
		},
	},
	{
		Key: "processes", Name: "获取进程信息", SQL: "select * from processes;",
		Columns: [][2]string{
			{"pid", "进程ID"},
			{"parent", "父进程ID"},
			{"name", "进程名"},
			{"path", "路径"},
			{"cmdline", "执行命令"},
			{"path", "路径"},
			{"start_time", "启动时间"},
		},
	},
	{
		Key: "listening_ports", Name: "获取监听端口信息", SQL: "select * from listening_ports;",
		Columns: [][2]string{
			{"address", "地址"},
			{"port", "端口号"},
			{"protocol", "协议"},
			{"pid", "进程ID"},
		},
	},
}

type osqueryAgentService struct{}

var OsqueryAgentService = new(osqueryAgentService)

func (s *osqueryAgentService) Query(q string) []*models.OsqueryAgent {
	ormer := orm.NewOrm()
	agents := make([]*models.OsqueryAgent, 0)
	queryset := ormer.QueryTable(new(models.OsqueryAgent))

	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)
	if q != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("name__icontains", q)
		cond = cond.AndCond(queryCond)
	}
	queryset.SetCond(cond).All(&agents)
	return agents
}

func (s *osqueryAgentService) GetByUUID(uuid string) *models.OsqueryAgent {
	ormer := orm.NewOrm()
	agent := &models.OsqueryAgent{
		UUID: uuid,
	}
	ormer.Read(agent, "UUID")
	return agent
}

type osqueryTaskService struct{}

var OsqueryTaskService = new(osqueryTaskService)

func (s *osqueryTaskService) Query(q string) []*models.OsqueryTask {
	ormer := orm.NewOrm()
	tasks := make([]*models.OsqueryTask, 0)
	queryset := ormer.QueryTable(new(models.OsqueryTask))

	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)
	if q != "" {
		queryCond := orm.NewCondition()
		queryCond = queryCond.Or("name__icontains", q)
		cond = cond.AndCond(queryCond)
	}
	queryset.SetCond(cond).OrderBy("-created_at").All(&tasks)
	return tasks
}

func (s *osqueryTaskService) Create(task *models.OsqueryTask, uuids []string) error {
	ormer := orm.NewOrm()
	var err error
	tx, err := ormer.Begin()
	if err != nil {
		return err
	}
	typ := s.GetTypeByKey(task.Type)
	if typ == nil {
		return fmt.Errorf("error type")
	}
	if _, err = tx.Insert(task); err == nil {
		for _, uuid := range uuids {
			job := &models.OsqueryJob{
				TaskId:  task.Id,
				UUID:    uuid,
				Content: typ.SQL,
				Status:  JobStatusNew,
				Result:  "",
			}
			_, err = tx.Insert(job)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}

func (s *osqueryTaskService) Delete(id int64) {

}

func (s *osqueryTaskService) Result(id int64) (*models.OsqueryTask, []*models.OsqueryJob, error) {
	task := &models.OsqueryTask{
		Id: id,
	}
	jobs := make([]*models.OsqueryJob, 0)

	ormer := orm.NewOrm()
	if err := ormer.Read(task); err != nil {
		return nil, nil, err
	}

	if _, err := ormer.QueryTable(new(models.OsqueryJob)).Filter("tid", id).All(&jobs); err != nil {
		return task, nil, err
	}
	return task, jobs, nil
}

func (s *osqueryTaskService) Types() []*models.OsqueryTaskType {
	return types
}

func (s *osqueryTaskService) GetTypeByKey(key string) *models.OsqueryTaskType {
	for _, typ := range types {
		if key == typ.Key {
			return typ
		}
	}
	return nil
}

func (s *osqueryTaskService) GetJobTotal(id int64) int64 {
	total, _ := orm.NewOrm().QueryTable(new(models.OsqueryJob)).Filter("tid", id).Count()
	return total
}

func (s *osqueryTaskService) GetJobComplated(id int64) int64 {
	total, _ := orm.NewOrm().QueryTable(new(models.OsqueryJob)).Filter("tid", id).Filter("status__in", []int{JobStatusSuccess, JobStatusFailure}).Count()
	return total
}

func (s *osqueryTaskService) GetJobRunning(id int64) int64 {
	total, _ := orm.NewOrm().QueryTable(new(models.OsqueryJob)).Filter("tid", id).Filter("status", JobStatusRunning).Count()
	return total
}
