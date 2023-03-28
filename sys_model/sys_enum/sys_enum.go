package sys_enum

import (
	sys_enum_audit "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/audit"
	sys_enum_auth "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/auth"
	pro_enum_business "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/business"
	sys_enum_casbin "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/casbin"
	sys_enum_file "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/file"
	sys_enum_license "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/license"
	sys_enum_organization "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/organization"
	sys_enum_oss "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/oss"
	sys_enum_permissions "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/permissions"
	sys_enum_role "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/role"
	sys_enum_upload "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/upload"
	sys_enum_user "github.com/SupenBysz/gf-admin-community/sys_model/sys_enum/internal/user"
)

type (
	BusinessType pro_enum_business.TypeEnum

	OssType sys_enum_oss.OssTypeEnum

	AuthActionType sys_enum_auth.ActionTypeEnum

	UploadEventState sys_enum_upload.EventStateEnum

	UserEvent sys_enum_user.EventEnum
	UserType  sys_enum_user.TypeEnum
	UserState sys_enum_user.StateEnum

	CabinEvent sys_enum_casbin.EventEnum

	PermissionMatchMode sys_enum_permissions.MatchModeEnum

	AuditAction = sys_enum_audit.ActionEnum
	AuditEvent  sys_enum_audit.EventEnum
)

var (
	Business = pro_enum_business.Business

	Oss    = sys_enum_oss.Oss
	Auth   = sys_enum_auth.Auth
	Upload = sys_enum_upload.Upload
	User   = sys_enum_user.User
	Casbin = sys_enum_casbin.Casbin
	File   = sys_enum_file.File

	Organization = sys_enum_organization.Organization
	Role         = sys_enum_role.Role
	Permissions  = sys_enum_permissions.Permissions

	Audit   = sys_enum_audit.Audit
	License = sys_enum_license.License
)
