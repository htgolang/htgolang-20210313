package forms

import (
	"cmdb/utils"

	"github.com/beego/beego/v2/core/validation"
)

type AssetForm struct {
	Id   int64  `form:"id"`
	Ip   string `form:"ip"`
	Addr string `form:"addr"`
}

func (f *AssetForm) Valid(v *validation.Validation) {
	v.Required(f.Ip, "ip.ip").Message("IP不能为空")
	v.Required(f.Addr, "addr.addr").Message("地址不能为空")
	v.IP(f.Ip, "ip.ip").Message("请输入正确的IP地址")

	if utils.CheckAssetExists(f.Id, f.Ip) {
		v.SetError("ip", "IP冲突")
	}
}
