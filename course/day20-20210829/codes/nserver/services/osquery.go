package services

import (
	"encoding/json"
	"strconv"
	"time"
)

const (
	JobStatusNew     = 0
	JobStatusRunning = 1
	JobStatusSuccess = 2
	JobStatusFailure = 3
)

type osqueryService struct {
}

var OsqueryService = new(osqueryService)

func (s *osqueryService) Register(uuid string, hostname string, addr string) error {
	// 插入数据库/更新数据
	var cnt int64
	err := Db.QueryRow("select count(*) as cnt from osquery_agent where uuid=?", uuid).Scan(&cnt)
	if err != nil {
		return err
	}

	now := time.Now()

	if cnt == 0 {
		// 新增
		_, err := Db.Exec("insert into osquery_agent(uuid, hostname, addr, created_at, updated_at, heartbeat_time)values(?, ?, ?, ?, ?, ?)", uuid, hostname, addr, now, now, now)
		return err
	} else {
		// 更新
		Db.Exec("update osquery_agent set hostname=?, addr=?, updated_at =? heartbeat_time=? where uuid=?", hostname, addr, now, now, uuid)
	}
	return nil
}

func (s *osqueryService) Heartbeat(uuid string) error {
	// 更新数据心跳时间
	now := time.Now()
	_, err := Db.Exec("update osquery_agent set heartbeat_time=? where uuid=?", now, uuid)
	return err
}

func (s *osqueryService) Config(id string, version string) (map[string]string, error) {
	// 查询返回配置
	result := map[string]string{
		"config": `
		{
			"schedule": {
				"check_etc_host": {
					"query": "select * from etc_hosts;",
					"interval" : 10
				}
			}
		}
		`,
	}

	return result, nil
}

func (s *osqueryService) Tasks(uuid string) (map[string]string, error) {
	rows, err := Db.Query("select id,content from osquery_job where uuid=? and status=?", uuid, JobStatusNew)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// key : id
	// value :sql
	result := map[string]string{}

	for rows.Next() {
		var (
			id      int64
			content string
		)
		if err := rows.Scan(&id, &content); err != nil {
			break
		}
		result[strconv.FormatInt(id, 10)] = content
		if _, err := Db.Exec("update osquery_job set status=? where id =?", JobStatusRunning, id); err != nil {
			break
		}
	}

	return result, nil
}

func (s *osqueryService) Result(uuid string, jobId int64, status int, result json.RawMessage) error {
	// 更新结果
	var jobStatus = JobStatusSuccess
	if status != 0 {
		jobStatus = JobStatusFailure
	}
	now := time.Now()
	_, err := Db.Exec("update osquery_job set status=?, result=?, complated_at =? where id=?", jobStatus, string(result), now, jobId)
	// 更新父任务状态
	// job 相同的tid 的jobs中存在状态为New, Running，则为运行中否则已完成

	var total int64
	if err := Db.QueryRow("select count(*) from osquery_job where tid=(select tid from osquery_job where id=?) and (status=? or status = ?)", jobId, JobStatusNew, JobStatusRunning).Scan(&total); err == nil && total == 0 {
		Db.Exec("update osquery_task set complated_at =? where id = (select tid from osquery_job where id=?)", now, jobId)
	}

	return err
}

func (s *osqueryService) Log(id string) error {
	// 插入日志
	return nil
}
