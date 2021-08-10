package controllers

import (
	"cmdb/forms"
	"cmdb/models"
	"cmdb/services"
	"log"
	"net/http"

	"github.com/beego/beego/v2/adapter/validation"
)

type AssetController struct {
	AuthenticationController
}

func (c *AssetController) Query() {
	keyword := c.GetString("q")
	assets := services.AssetService.QueryAsset(keyword)
	c.Data["Assets"] = assets
	c.TplName = "asset/list.html"
}

func (c *AssetController) Delete() {
	id, _ := c.GetInt64("id")
	services.AssetService.DleteAsset(id)
	c.Redirect("/asset/query/", http.StatusFound)
}

func (c *AssetController) Add() {
	var form forms.AssetForm
	var valid validation.Validation

	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(&form)
		if err == nil {
			if b, err := valid.Valid(&form); err != nil {
				log.Println("valid add asset data error.", err)
				valid.SetError("asset", "验证数据错误")
			} else if b {
				err := services.AssetService.AddAsset(form.Asset)
				if err == nil {
					c.Redirect("/asset/query/", http.StatusFound)
					return
				}
				log.Println("add asset error.", err)
				valid.SetError("asset", "添加资产失败")
			}
		} else {
			log.Println("parse add asset data error.", err)
			valid.SetError("asset", "提交数据错误")
		}
	}
	c.Data["ErrMsgs"] = valid.ErrorMap()
	c.TplName = "asset/add.html"
}

func (c *AssetController) Edit() {
	var form forms.AssetForm
	var valid validation.Validation
	var asset *models.Asset = &models.Asset{}

	if c.Ctx.Input.IsGet() {
		id, _ := c.GetInt64("id")
		asset = services.AssetService.QueryAssetByID(id)
	}

	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(&form)
		if err == nil {
			if b, err := valid.Valid(&form); err != nil {
				log.Println("valid edit asset data error.", err)
				valid.SetError("asset", "验证数据错误")
			} else if b {
				err := services.AssetService.EditAsset(form.Asset)
				if err == nil {
					c.Redirect("/asset/query/", http.StatusFound)
					return
				}
				log.Println("edit asset error.", err)
				valid.SetError("asset", "修改资产信息失败")
			} else {
				asset.Id = form.Id
				asset.Ip = form.Ip
				asset.Addr = form.Addr
			}
		} else {
			log.Println("parse edit form data error.", err)
			valid.SetError("asset", "提交数据错误")
		}
	}
	c.Data["ErrMsgs"] = valid.ErrorMap()
	c.Data["Asset"] = asset
	c.TplName = "asset/edit.html"
}
