package services

import (
	"github.com/beego/beego/v2/client/orm"

	"cmdb/models"
)

const (
	keyScan = "asset:scan"
)

type configService struct{}

var ConfigService = new(configService)

func (s *configService) GetAssetScan() *models.Config {
	ormer := orm.NewOrm()
	config := &models.Config{Key: keyScan}
	if err := ormer.Read(config, "Key"); err == nil {
		return config
	}
	return nil
}

func (s *configService) ModifyAssetScan(config *models.Config) {
	ormer := orm.NewOrm()
	tempconfig := &models.Config{Key: keyScan}
	if err := ormer.Read(tempconfig, "Key"); err != nil {
		// 添加
		config.Key = keyScan
		ormer.Insert(config)
	} else {
		tempconfig.Value = config.Value
		ormer.Update(tempconfig)
	}
}
