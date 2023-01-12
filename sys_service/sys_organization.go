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
	ISysOrganization interface {
		QueryOrganizationList(ctx context.Context, info sys_model.SearchParams) (*sys_model.OrganizationInfoListRes, error)
		GetOrganizationList(ctx context.Context, parentId int64, IsRecursive bool) ([]*sys_entity.SysOrganization, int, error)
		GetOrganizationTree(ctx context.Context, parentId int64) ([]*sys_model.SysOrganizationTree, error)
		CreateOrganizationInfo(ctx context.Context, info sys_model.SysOrganizationInfo) (*sys_entity.SysOrganization, error)
		UpdateOrganizationInfo(ctx context.Context, info sys_model.SysOrganizationInfo) (*sys_entity.SysOrganization, error)
		SaveOrganizationInfo(ctx context.Context, info sys_model.SysOrganizationInfo) (*sys_entity.SysOrganization, error)
		GetOrganizationInfo(ctx context.Context, id int64) (*sys_entity.SysOrganization, error)
		DeleteOrganizationInfo(ctx context.Context, id int64) (bool, error)
	}
)

var (
	localSysOrganization ISysOrganization
)

func SysOrganization() ISysOrganization {
	if localSysOrganization == nil {
		panic("implement not found for interface ISysOrganization, forgot register?")
	}
	return localSysOrganization
}

func RegisterSysOrganization(i ISysOrganization) {
	localSysOrganization = i
}
