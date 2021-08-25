package services

import (
	"cmdb/models"
	"log"

	"github.com/beego/beego/v2/adapter/orm"
)

type moduleService struct{}

var ModuleService = moduleService{}

func (s *moduleService) QueryModuleByUrl(url string) *models.Module {
	// fmt.Println(url)
	ormer := orm.NewOrm()
	module := models.Module{Url: url}
	if err := ormer.Read(&module, "url"); err != nil {
		log.Println("Failed to query module by url.", err)
		return nil
	}
	return &module
}
