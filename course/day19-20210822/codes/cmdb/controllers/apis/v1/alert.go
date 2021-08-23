package v1

import (
	"encoding/json"

	"cmdb/controllers/apis"
	"cmdb/forms"
	"cmdb/responses"
	"cmdb/services"
)

type AlertController struct {
	apis.ApiController
}

func (c *AlertController) Trigger() {
	c.Ctx.Input.CopyBody(1024 * 1024)

	var form forms.PrometheusAlertsForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		for _, alert := range form.Alerts {
			services.PrometheusAlertService.Trigger(alert.ToModel())
		}
	}
	c.Data["json"] = responses.NewJsonSuccess()
	c.ServeJSON()
}
