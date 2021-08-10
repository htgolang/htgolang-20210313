package forms

import (
	"cmdb/models"
	"cmdb/utils"

	"github.com/beego/beego/v2/core/validation"
)

type AssetForm struct {
	Id    int64  `form:"id"`
	Ip    string `form:"ip"`
	Addr  string `form:"addr"`
	Asset *models.Asset
}

func (f *AssetForm) Valid(v *validation.Validation) {
	v.Required(f.Ip, "ip.ip").Message("IP不能为空")
	v.Required(f.Addr, "addr.addr").Message("地址不能为空")
	v.MaxSize(f.Addr, 128, "addr.addr").Message("地址不超过128个字符")
	v.IP(f.Ip, "ip.ip").Message("请输入正确的IP地址")

	if utils.CheckAssetExists(f.Id, f.Ip) {
		v.SetError("ip", "IP冲突")
	}

	if !v.HasErrors() {
		f.Asset = new(models.Asset)
		f.Asset.Id = f.Id
		f.Asset.Ip = f.Ip
		f.Asset.Addr = f.Addr
	}
}
