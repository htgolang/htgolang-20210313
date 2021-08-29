package controllers

import (
	"net/http"
	"strings"

	"cmdb/forms"
	"cmdb/services"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

type OsqueryTaskController struct {
	LayoutController
}

func (c *OsqueryTaskController) Query() {
	q := strings.TrimSpace(c.GetString("q"))
	c.Data["objects"] = services.OsqueryTaskService.Query(q)
	c.Data["q"] = q
	c.Data["osqueryTaskService"] = services.OsqueryTaskService
	c.TplName = "osquery/task/query.html"
}

func (c *OsqueryTaskController) Create() {
	form := new(forms.OsqueryTaskForm)
	valid := new(validation.Validation)
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("default", "提交数据错误")
			} else if ok {
				// 存储数据
				services.OsqueryTaskService.Create(form.ToModel(), form.UUIDs) // 根据存储结果给用户响应

				c.Redirect(web.URLFor("OsqueryTaskController.Query"), http.StatusFound)
			}
		} else {
			valid.SetError("default", "提交数据错误")
		}
	}

	c.Data["form"] = form
	c.Data["errors"] = valid.ErrorsMap
	c.Data["types"] = services.OsqueryTaskService.Types()
	c.Data["targets"] = services.OsqueryAgentService.Query("")
	c.TplName = "osquery/task/create.html"
}

func (c *OsqueryTaskController) Delete() {

}

func (c *OsqueryTaskController) Result() {
	if id, err := c.GetInt64("id"); err == nil {
		task, jobs, _ := services.OsqueryTaskService.Result(id)
		c.Data["task"] = task
		c.Data["jobs"] = jobs
		c.Data["taskType"] = services.OsqueryTaskService.GetTypeByKey(task.Type)
		c.Data["osqueryTaskService"] = services.OsqueryTaskService
		c.Data["osqueryAgentService"] = services.OsqueryAgentService
	}
	c.TplName = "osquery/task/result.html"
}
