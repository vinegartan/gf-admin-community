package menu

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

type PermissionTypeEnum = *sys_model.SysPermissionTree

type permissionType struct {
	ViewDetail PermissionTypeEnum
	Tree       PermissionTypeEnum
	Update     PermissionTypeEnum
	Delete     PermissionTypeEnum
	Create     PermissionTypeEnum
}

var PermissionType = permissionType{
	ViewDetail: permission.New(5948649434487532, "ViewDetail", "查看菜单", "查看某个菜单"),
	Tree:       permission.New(5948649530309523, "Tree", "菜单树", "查看菜单树"),
	Update:     permission.New(5948649642787342, "Update", "更新菜单", "更新某个菜单"),
	Delete:     permission.New(5948649739509856, "Delete", "删除菜单", "删除某个菜单"),
	Create:     permission.New(5948649828767321, "Create", "创建菜单", "创建菜单"),
}
