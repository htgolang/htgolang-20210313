package services

import (
	"cmdb/models"
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type roleService struct{}

var RoleService roleService

func (s *roleService) QueryRole() []models.Role {
	var roles []models.Role
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(models.Role))

	if _, err := queryset.All(&roles); err != nil {
		log.Println("query roles error.", err)
	}
	return roles
}

func (s *roleService) QueryRoleByID(id int64) *models.Role {
	ormer := orm.NewOrm()
	// fmt.Println(id)
	var role = models.Role{Id: id}
	err := ormer.Read(&role)
	if err != nil {
		log.Println("query role by id error.", err)
		return nil
	}
	return &role
}

func (s *roleService) GetModules(id int64) []models.Module {
	ormer := orm.NewOrm()

	//查询role关联的module id
	res := make([]models.RelRoleModule, 0)
	moduleIds := make([]int, 0)
	queryset := ormer.QueryTable(new(models.RelRoleModule)).Filter("role_id__exact", id)
	if _, err := queryset.All(&res, "module_id"); err != nil {
		log.Println("query the relationship between roles and modules faild.", err)
	} else {
		for _, v := range res {
			moduleIds = append(moduleIds, int(v.ModuleId))
		}
	}

	//查询role对应的module
	modules := make([]models.Module, 0)
	moduleQueryset := ormer.QueryTable(new(models.Module)).Filter("id__in", moduleIds)
	if _, err := moduleQueryset.All(&modules); err != nil {
		log.Println(" Failed to query the modules associated with the role.", err)
	}

	return modules
}

func (s *roleService) IsPermission(roleid int64, url string) bool {
	modules := s.GetModules(roleid)
	for _, v := range modules {
		if url == v.Url {
			return true
		}
	}
	return false
}
