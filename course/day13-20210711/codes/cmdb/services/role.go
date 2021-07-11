package services

import (
	"cmdb/models"
)

var RoleService = new(roleService)

var modules = []*models.Module{
	{Name: "资产管理", Path: "/asset/query", Controller: "AssetController", Action: "Query", IsMenu: true},
	{Name: "告警管理", Path: "/alarm/query", Controller: "AlarmController", Action: "Query", IsMenu: true},
	{Name: "部门管理", Path: "/department/query", Controller: "DepartmentController", Action: "Query", IsMenu: true},
	{Name: "用户管理", Path: "/user/query", Controller: "UserController", Action: "Query", IsMenu: true},
	{Name: "添加用户管理", Path: "/user/create", Controller: "UserController", Action: "Create", IsMenu: false},
	{Name: "修改用户管理", Path: "/user/modify", Controller: "UserController", Action: "Modify", IsMenu: false},
	{Name: "删除用户管理", Path: "/user/delete", Controller: "UserController", Action: "Delete", IsMenu: false},
	{Name: "修改自己的密码", Path: "/user/modifypassword", Controller: "UserController", Action: "ModifyPassword", IsMenu: false},
	{Name: "角色管理", Path: "/role/query", Controller: "RoleController", Action: "Query", IsMenu: true},
}

type roleService struct {
}

func (s *roleService) GetMenus(roleId int64) []*models.Module {
	menus := make([]*models.Module, 0, 10)
	for _, module := range modules {
		// roleId 过滤由权限的菜单

		if module.IsMenu {
			menus = append(menus, module)
		}
	}
	return menus
}

func (s *roleService) GetModules(roleId int64) []*models.Module {
	permissionModules := make([]*models.Module, 0, 10)
	for _, module := range modules {
		// roleId 过滤由权限的菜单
		permissionModules = append(permissionModules, module)
	}
	return permissionModules
}

func (s *roleService) IsPermission(roleId int64, controller, action string) bool {
	modules := s.GetModules(roleId)
	for _, module := range modules {
		if module.Controller == controller && module.Action == action {
			return true
		}
	}
	return false
}
