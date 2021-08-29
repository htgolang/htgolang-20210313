package forms

import (
	"cmdb/models"
	"encoding/json"

	"github.com/beego/beego/v2/core/validation"
)

type AssetScanConfigForm struct {
	Enable   bool   `form:"enable" json:"enable"`
	Addr     string `form:"addr" json:"addr"`
	Interval int    `form:"interval" json:"interval"`
	Ports    string `form:"ports" json:"ports"`
}

func (f *AssetScanConfigForm) Valid(v *validation.Validation) {

}

func (f *AssetScanConfigForm) FillModel(config *models.Config) {
	if config == nil || json.Unmarshal([]byte(config.Value), f) != nil {
		// 若未配置,反序列化失败 则使用默认值
		f.Enable = false
		f.Addr = ""
		f.Interval = 300
		f.Ports = ""
	}
}

func (f *AssetScanConfigForm) ToModel() *models.Config {
	config := new(models.Config)

	if value, err := json.Marshal(f); err == nil {
		config.Value = string(value)
	}

	return config
}
