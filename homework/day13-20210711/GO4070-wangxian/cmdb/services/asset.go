package services

import (
	"cmdb/models"
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type assetService struct{}

var AssetService assetService

func (s *assetService) QueryAsset(keyword string) []models.Asset {
	var assets []models.Asset
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.Asset))

	if keyword != "" {
		cond := orm.NewCondition().And("ip__contains", keyword).Or("addr__contains", keyword)
		queryset = queryset.SetCond(cond)
	}

	if _, err := queryset.All(&assets); err != nil {
		log.Println("query asset error.", err)
		return nil
	}
	return assets

}

func (s *assetService) DleteAsset(id int64) {
	var asset = models.Asset{Id: id}
	ormer := orm.NewOrm()

	if _, err := ormer.Delete(&asset); err != nil {
		log.Println("delete asset error.", err)
	}
}

func (s *assetService) AddAsset(asset *models.Asset) error {
	ormer := orm.NewOrm()
	if _, err := ormer.Insert(asset); err != nil {
		log.Println("add asset error.", err)
		return err
	}
	return nil
}

func (s *assetService) EditAsset(asset *models.Asset) error {
	ormer := orm.NewOrm()

	if _, err := ormer.Update(asset); err != nil {
		log.Println("edit asset error.", err)
		return err
	}
	return nil
}

func (s *assetService) QueryAssetByID(id int64) *models.Asset {
	var asset = models.Asset{Id: id}
	ormer := orm.NewOrm()
	if err := ormer.Read(&asset); err != nil {
		log.Println("query asset by id error.", err)
		return nil
	}
	return &asset
}
