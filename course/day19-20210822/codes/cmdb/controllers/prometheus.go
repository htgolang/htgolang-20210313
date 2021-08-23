package controllers

import (
	"cmdb/forms"
	"cmdb/services"
	"net/http"

	"github.com/beego/beego/v2/core/utils/pagination"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

type PrometheusAgentController struct {
	LayoutController
}

func (c *PrometheusAgentController) Query() {
	q := c.GetString("q")

	c.Data["objects"] = services.PrometheusAgentService.Query(q)
	c.Data["q"] = q

	c.TplName = "prometheus/agent/query.html"
}

func (c *PrometheusAgentController) Modify() {
	form := &forms.ModifyAgentForm{}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsGet() {
		id, _ := c.GetInt64("id")
		agent := services.PrometheusAgentService.GetById(id)
		if agent == nil {
			// 存在错误
			valid.SetError("default", "修改对象不存在")
		} else {
			form.FillModel(agent)
		}
	} else {
		if err := c.ParseForm(form); err != nil {
			valid.SetError("default", "提交数据错误")
		} else if ok, err := valid.Valid(form); err != nil {
			valid.SetError("default", "提交数据错误")
		} else if ok {
			// 成功
			services.PrometheusAgentService.Modify(form.ToModel())
			c.Redirect(web.URLFor("PrometheusAgentController.Query"), http.StatusFound)
			return
		}
	}

	c.Data["form"] = form
	c.Data["errors"] = valid.ErrorsMap
	c.TplName = "prometheus/agent/modify.html"
	c.LayoutSections["SectionStyles"] = "prometheus/agent/modify_style.html"
	c.LayoutSections["SectionScripts"] = "prometheus/agent/modify_script.html"
}

func (c *PrometheusAgentController) Delete() {
	if id, err := c.GetInt64("id"); err == nil {
		services.PrometheusAgentService.Delete(id)
	}
	c.Redirect(web.URLFor("PrometheusAgentController.Query"), http.StatusFound)
}

type PrometheusAlertController struct {
	LayoutController
}

func (c *PrometheusAlertController) Query() {
	q := c.GetString("q")
	var pageSize int = 10

	total := services.PrometheusAlertService.Count(q)
	paginator := pagination.NewPaginator(c.Ctx.Request, pageSize, total)
	// offset limit

	c.Data["objects"] = services.PrometheusAlertService.Query(q, paginator.Offset(), paginator.PerPageNums)
	c.Data["q"] = q
	c.Data["paginator"] = paginator

	c.TplName = "prometheus/alert/query.html"
}
