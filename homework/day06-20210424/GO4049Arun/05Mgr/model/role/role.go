package role

type Role struct {
	Admin int
	Manager int
	Enginner int
}
var roleConst = Role{Admin: 0, Manager: 1, Enginner: 2}

func GetRole() {

}

func init() {

}
//todo:基于角色的权限认证先保留;


