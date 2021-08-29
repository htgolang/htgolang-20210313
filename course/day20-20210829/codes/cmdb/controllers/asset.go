package controllers

import (
	"cmdb/schedulers"
	"cmdb/services"
	"net/http"
	"strings"

	"cmdb/forms"
	"cmdb/utils"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

type AssetController struct {
	LayoutController
}

func (c *AssetController) Config() {
	form := new(forms.AssetScanConfigForm)
	valid := new(validation.Validation)
	success := ""
	// get获取
	// post更新
	if c.Ctx.Input.IsPost() {
		// Post请求 创建用户
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("default", "提交数据错误")
			} else if ok {
				// 存储数据
				services.ConfigService.ModifyAssetScan(form.ToModel()) // 根据存储结果给用户响应

				schedulers.AssetScheduler.SyncConfig() // 同步调度配置
				success = "配置成功"
			}
		} else {
			valid.SetError("default", "提交数据错误")
		}
	} else {
		config := services.ConfigService.GetAssetScan()
		form.FillModel(config)
	}

	c.Data["form"] = form
	c.Data["errors"] = valid.ErrorMap()
	c.Data["success"] = success
	c.TplName = "asset/scan_config.html"
}

func (c *AssetController) Scan() {
	form := new(forms.AssetScanConfigForm)
	config := services.ConfigService.GetAssetScan()
	form.FillModel(config)

	ipRanges := utils.ParseIPRange(form.Addr)
	ports := utils.ParsePorts(form.Ports)

	services.AssetService.Scan(ipRanges, ports)
	c.Redirect(web.URLFor("AssetController.Query"), http.StatusFound)
}

func (c *AssetController) Query() {
	q := strings.TrimSpace(c.GetString("q"))
	c.Data["objects"] = services.AssetService.Query(q)
	c.Data["q"] = q
	c.Data["scanning"] = services.AssetService.IsScanning()
	c.TplName = "asset/query.html"
}
