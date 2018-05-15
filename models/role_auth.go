package models

type RoleAuth struct {
	RoleId uint `orm:"column(role_id)" description:"角色ID"`
	AuthId uint `orm:"column(auth_id)" description:"权限ID"`
}
