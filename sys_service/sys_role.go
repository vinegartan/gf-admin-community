// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
)

type (
	ISysRole interface {
		QueryRoleList(ctx context.Context, info sys_model.SearchParams, unionMainId int64) (*sys_model.RoleListRes, error)
		GetRoleById(ctx context.Context, id int64) (*sys_entity.SysRole, error)
		Create(ctx context.Context, info sys_model.SysRole) (*sys_entity.SysRole, error)
		Update(ctx context.Context, info sys_model.SysRole) (*sys_entity.SysRole, error)
		Save(ctx context.Context, info sys_model.SysRole) (*sys_entity.SysRole, error)
		Delete(ctx context.Context, roleId int64) (bool, error)
		SetRoleForUser(ctx context.Context, roleId, userId, makeUserUnionMainId int64) (bool, error)
		RemoveRoleForUser(ctx context.Context, roleId int64, userId int64) (bool, error)
		GetRoleUsers(ctx context.Context, roleId int64, makeUserUnionMainId int64) (*[]sys_model.SysUser, error)
		GetUserRoleList(ctx context.Context, userId int64) (*[]sys_entity.SysRole, error)
		SetRolePermissions(ctx context.Context, roleId int64, permissionIds []int64, makeUserUnionMainId int64) (bool, error)
		GetRolePermissions(ctx context.Context, roleId int64) ([]int64, error)
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
