// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysAudit interface {
		InstallHook(state sys_enum.AuditEvent, category int, hookFunc sys_hook.AuditHookFunc) int64
		UnInstallHook(savedHookId int64)
		CleanAllHook()
		QueryAuditList(ctx context.Context, filter *base_model.SearchParams) (*sys_model.AuditListRes, error)
		GetAuditById(ctx context.Context, id int64) *sys_entity.SysAudit
		GetAuditByLatestUnionMainId(ctx context.Context, unionMainId int64) *sys_entity.SysAudit
		CreateAudit(ctx context.Context, info sys_model.CreateSysAudit) (*sys_entity.SysAudit, error)
		UpdateAudit(ctx context.Context, id int64, state int, reply string, auditUserId int64) (bool, error)
	}
)

var (
	localSysAudit ISysAudit
)

func SysAudit() ISysAudit {
	if localSysAudit == nil {
		panic("implement not found for interface ISysAudit, forgot register?")
	}
	return localSysAudit
}

func RegisterSysAudit(i ISysAudit) {
	localSysAudit = i
}
