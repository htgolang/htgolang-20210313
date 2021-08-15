package v1

import (
	"encoding/json"

	"cmdb/controllers/apis"
	"cmdb/forms"
	"cmdb/responses"
	"cmdb/services"
)

type AgentController struct {
	apis.ApiController
}

func (c *AgentController) Heartbeat() {
	c.Ctx.Input.CopyBody(1024 * 1024)

	var form forms.AgentdHeartbeatForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		services.PrometheusAgentService.Heartbeat(form.UUID)
		c.Data["json"] = responses.NewJsonSuccess()
	} else {
		c.Data["json"] = responses.NewJsonFailure(err.Error())
	}

	c.ServeJSON()
}

func (c *AgentController) Register() {
	c.Ctx.Input.CopyBody(1024 * 1024)

	var form forms.AgentdRegisterForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		services.PrometheusAgentService.Register(form.UUID, form.Addr, form.Hostname)
		c.Data["json"] = responses.NewJsonSuccess()
	} else {
		c.Data["json"] = responses.NewJsonFailure(err.Error())
	}
	c.ServeJSON()
}

func (c *AgentController) Config() {
	uuid := c.GetString("uuid")
	version, _ := c.GetInt64("version", 0)
	if config, version, err := services.PrometheusAgentService.GetConfigByUUID(uuid, version); err != nil {
		// 错误
		c.Data["json"] = responses.NewJsonFailure(err.Error())
	} else if config == "" {
		// 不需要更新
		c.Data["json"] = responses.NewJsonResponse(struct {
			Updated bool   `json:"updated"`
			Version int64  `json:"version"`
			Config  string `json:"config"`
		}{false, version, ""})
	} else {
		c.Data["json"] = responses.NewJsonResponse(struct {
			Updated bool   `json:"updated"`
			Version int64  `json:"version"`
			Config  string `json:"config"`
		}{true, version, config})
	}

	c.ServeJSON()
}
